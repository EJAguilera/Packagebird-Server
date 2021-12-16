package main

import (
	protobuf_grpc "Server/FileTransfer"
	packageops_grpc "Server/PackageOperations"
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

const (
	ADDRESS     = "127.0.0.1"
	PORT        = 50051
	BUFFER_SIZE = 64 * 1024
)

type Server struct {
	protobuf_grpc.UnimplementedFileServiceServer
	packageops_grpc.UnimplementedPackageOperationServicesServer
}

func (s *Server) ListPackages(ctx context.Context, request *packageops_grpc.PackageListRequest) (*packageops_grpc.PackageList, error) {
	packages, err := ioutil.ReadDir("packages/")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("List Packages Request Received...\n")

	var contents string
	for _, packagefile := range packages {
		contents += fmt.Sprintf("%v\n", packagefile.Name())
	}

	packagelist := &packageops_grpc.PackageList{
		List: contents,
	}

	return packagelist, nil
}

func (s *Server) TestPackage(ctx context.Context, request *packageops_grpc.PackageTestRequest) (*packageops_grpc.PackageTestResponse, error) {
	packagename := fmt.Sprintf("packages/%v.tar.gz", request.GetPackagename())
	if _, err := os.Stat(packagename); err == nil {
		server_message := fmt.Sprintf("Package %v running test...", request.GetPackagename())
		fmt.Println(server_message)
		time.Sleep(2 * time.Second)
		server_message = fmt.Sprintf("Package %v passed test...", request.GetPackagename())
		response := &packageops_grpc.PackageTestResponse{
			Response: server_message,
		}
		return response, nil

	} else if errors.Is(err, os.ErrNotExist) {
		message := fmt.Sprintf("Package %v not found!", request.GetPackagename())
		response := &packageops_grpc.PackageTestResponse{
			Response: message,
		}
		return response, nil
	} else {
		message := "Unknown Error Encountered!"
		response := &packageops_grpc.PackageTestResponse{
			Response: message,
		}
		return response, nil
	}
}

func (s *Server) BuildPackage(ctx context.Context, request *packageops_grpc.PackageBuildRequest) (*packageops_grpc.PackageBuildResponse, error) {
	packagename := fmt.Sprintf("packages/%v.tar.gz", request.GetPackagename())
	if _, err := os.Stat(packagename); err == nil {
		server_message := fmt.Sprintf("Package %v being built...", request.GetPackagename())
		fmt.Println(server_message)
		time.Sleep(2 * time.Second)
		server_message = fmt.Sprintf("Package %v built...", request.GetPackagename())
		response := &packageops_grpc.PackageBuildResponse{
			Response: server_message,
		}
		return response, nil

	} else if errors.Is(err, os.ErrNotExist) {
		message := fmt.Sprintf("Package %v not found!", request.GetPackagename())
		response := &packageops_grpc.PackageBuildResponse{
			Response: message,
		}
		return response, nil
	} else {
		message := "Unknown Error Encountered!"
		response := &packageops_grpc.PackageBuildResponse{
			Response: message,
		}
		return response, nil
	}
}

func FindPackage(packagename string) bool {
	return true
}

// RPC for downloading files from server to client
func (s *Server) Download(request *protobuf_grpc.Request, fileStream protobuf_grpc.FileService_DownloadServer) error {
	fmt.Printf("Recieved request for downloading package %v", request.GetBody())

	package_path := fmt.Sprintf("packages/%v", request.GetBody())

	file, err := os.Open(package_path)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	buffer := make([]byte, BUFFER_SIZE)

	for {
		bytesRead, err := file.Read(buffer)

		// Oddly written way of intercepting EOF
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}

		file := &protobuf_grpc.File{
			Content: &protobuf_grpc.File_Chunk{
				Chunk: buffer[:bytesRead],
			},
		}

		err = fileStream.Send(file)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	fmt.Println("Successfully completed download operation.")
	return nil
}

// RPC for uploading files from client to server
func (s *Server) Upload(fileStream protobuf_grpc.FileService_UploadServer) error {

	fmt.Println("Received request to upload file.")
	file, err := os.Create("tmp/temp.bin")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	for {
		chunk, err := fileStream.Recv()
		// Chunk is empty
		if (chunk == nil) || (len(chunk.GetChunk()) == 0) {
			break
		}

		// If error is found
		if err != nil {
			log.Fatal(err)
			return err
		}

		// Otherwise, write file in
		_, err = file.Write(chunk.GetChunk())
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
			return err
		}
	}

	message := &protobuf_grpc.Response{
		Body: "File uploaded successfully!",
	}

	fileStream.SendAndClose(message)
	fmt.Println("File upload operation completed successfully.")
	return nil
}

// Names file uploaded in a sequential request
func (s *Server) NameFile(context context.Context, request *protobuf_grpc.Request) (*protobuf_grpc.Response, error) {
	oldFileName := "tmp/temp.bin"
	newFileName := fmt.Sprintf("packages/%v", request.GetBody())

	err := os.Rename(oldFileName, newFileName)
	if err != nil {
		log.Fatal(err)

		response := &protobuf_grpc.Response{
			Header: "File Rename",
			Body:   "Failed to rename file!",
		}

		return response, err
	}

	response := &protobuf_grpc.Response{
		Header: "File Rename",
		Body:   "Succesfully renamed file!",
	}

	return response, nil
}

func GetDatabase(path string) (bool, error) {
	// Get package database from path
	_, err := os.Open(path)

	// If package-database does not exist,
	if errors.Is(err, os.ErrNotExist) {
		// Greate package database
		_, err := os.Create("packages.db")
		if err != nil {
			message := fmt.Sprintf("Error creating package database file! Error: %v", err)
			fmt.Println(message)
			return false, err
		} else {
			return true, nil
		}
	} else {
		// Database file found
		return true, nil
	}
}

type PackageRegistry struct {
	registry *sql.DB
}

func NewPackageRegistry(registry *sql.DB) *PackageRegistry {
	return &PackageRegistry{
		registry: registry,
	}
}

func FormatPackageRegistry(packages *PackageRegistry) error {
	query := `
	CREATE TABLE IF NOT EXISTS packages(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		version INTEGER,
		path TEXT NOT NULL
	);
	`
	_, err := packages.registry.Exec(query)
	return err
}

func CreatePackage(packages *PackageRegistry, name string, description string, version int, path string) error {
	query := `
	INSERT INTO packages (name, description, version, path) VALUES(?,?,?,?);
	`
	_, err := packages.registry.Exec(query, name, description, version, path)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

type packageType struct {
	name        string
	description string
	version     int
	path        string
}

func ListPackages(packages *PackageRegistry) error {
	query := `
	SELECT name, description, version, path FROM packages
	`

	response, err := packages.registry.Query(query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer response.Close()

	var allRows []packageType

	for response.Next() {
		var row packageType
		if err := response.Scan(&row.name, &row.description, &row.version, &row.path); err != nil {
			return err
		}
		allRows = append(allRows, row)
	}

	for index := range allRows {
		fmt.Println(fmt.Sprintf("Name: %v, Description: %v, Version: %v, Path: %v", allRows[index].name, allRows[index].description, allRows[index].version, allRows[index].path))
	}

	return nil
}

// Main function for listening to incoming request from clients
func main() {
	flag.Parse()
	// Get or create Packages database
	// GetDatabase("packages.db")
	// Connect to Packages database
	db, err := sql.Open("sqlite3", "packages.db")
	if err != nil {
		log.Fatal(err)
	}
	package_registry := NewPackageRegistry(db)
	FormatPackageRegistry(package_registry)
	// CreatePackage(package_registry, "Markle", "Markup-related functions.", 1, "./packages/markle-v1.tar.gz")
	// ListPackages(package_registry)

	listener, err := net.Listen("tcp", fmt.Sprintf("%v:%v", ADDRESS, PORT))
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}

	server := grpc.NewServer()
	protobuf_grpc.RegisterFileServiceServer(server, &Server{})
	packageops_grpc.RegisterPackageOperationServicesServer(server, &Server{})
	log.Printf("Server listening at address %v", listener.Addr())

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Error serving: %v", err)
	}
}
