import grpc
import FileTransfer_pb2
import FileTransfer_pb2_grpc

# Parses content of file
def file_reader(filename):
    BUFFER_SIZE = 64 * 1024
    with open(filename, 'rb') as file:
        while True:
            chunks = file.read(BUFFER_SIZE)
            if len(chunks) == 0:
                return
            yield FileTransfer_pb2.File(chunk=chunks)

# Downloads file from server
def download(filename):
    with grpc.insecure_channel('127.0.0.1:50051') as channel:
        stub = FileTransfer_pb2_grpc.FileServiceStub(channel)
        response = stub.Download(FileTransfer_pb2.Request(body=f'{filename}'))
        with open(filename, 'wb') as file:
            for chunk in response:
                file.write(chunk.chunk)
        print(f'Downloaded file: {filename}')

# Uploads file to server
def upload(filename):
    with grpc.insecure_channel('127.0.0.1:50051') as channel:
        stub = FileTransfer_pb2_grpc.FileServiceStub(channel)
        file_chunks = file_reader(filename)
        response = stub.Upload(file_chunks)
        print(f'Uploaded file: {filename}, with response: {response}')
        nameFileRequest = FileTransfer_pb2.Request(header='Rename File', body=f'{filename}')
        response = stub.NameFile(nameFileRequest)
        print(f'Successfully renamed remote temp file to {filename},\nresponse: {response}')

# For test calling
if __name__ == '__main__':
    download('Test.txt')
    upload('test.tar.gz')