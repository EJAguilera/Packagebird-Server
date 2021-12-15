# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: PackageOperations.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='PackageOperations.proto',
  package='PackageOperations',
  syntax='proto3',
  serialized_options=b'Z\023./PackageOperations',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x17PackageOperations.proto\x12\x11PackageOperations\"\x14\n\x12PackageListRequest\"\x1b\n\x0bPackageList\x12\x0c\n\x04list\x18\x01 \x01(\t\")\n\x12PackageTestRequest\x12\x13\n\x0bpackagename\x18\x01 \x01(\t\"\'\n\x13PackageTestResponse\x12\x10\n\x08response\x18\x01 \x01(\t\"*\n\x13PackageBuildRequest\x12\x13\n\x0bpackagename\x18\x01 \x01(\t\"(\n\x14PackageBuildResponse\x12\x10\n\x08response\x18\x01 \x01(\t2\xb0\x02\n\x18PackageOperationServices\x12U\n\x0cListPackages\x12%.PackageOperations.PackageListRequest\x1a\x1e.PackageOperations.PackageList\x12\\\n\x0bTestPackage\x12%.PackageOperations.PackageTestRequest\x1a&.PackageOperations.PackageTestResponse\x12_\n\x0c\x42uildPackage\x12&.PackageOperations.PackageBuildRequest\x1a\'.PackageOperations.PackageBuildResponseB\x15Z\x13./PackageOperationsb\x06proto3'
)




_PACKAGELISTREQUEST = _descriptor.Descriptor(
  name='PackageListRequest',
  full_name='PackageOperations.PackageListRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=46,
  serialized_end=66,
)


_PACKAGELIST = _descriptor.Descriptor(
  name='PackageList',
  full_name='PackageOperations.PackageList',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='list', full_name='PackageOperations.PackageList.list', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=68,
  serialized_end=95,
)


_PACKAGETESTREQUEST = _descriptor.Descriptor(
  name='PackageTestRequest',
  full_name='PackageOperations.PackageTestRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='packagename', full_name='PackageOperations.PackageTestRequest.packagename', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=97,
  serialized_end=138,
)


_PACKAGETESTRESPONSE = _descriptor.Descriptor(
  name='PackageTestResponse',
  full_name='PackageOperations.PackageTestResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='response', full_name='PackageOperations.PackageTestResponse.response', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=140,
  serialized_end=179,
)


_PACKAGEBUILDREQUEST = _descriptor.Descriptor(
  name='PackageBuildRequest',
  full_name='PackageOperations.PackageBuildRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='packagename', full_name='PackageOperations.PackageBuildRequest.packagename', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=181,
  serialized_end=223,
)


_PACKAGEBUILDRESPONSE = _descriptor.Descriptor(
  name='PackageBuildResponse',
  full_name='PackageOperations.PackageBuildResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='response', full_name='PackageOperations.PackageBuildResponse.response', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=225,
  serialized_end=265,
)

DESCRIPTOR.message_types_by_name['PackageListRequest'] = _PACKAGELISTREQUEST
DESCRIPTOR.message_types_by_name['PackageList'] = _PACKAGELIST
DESCRIPTOR.message_types_by_name['PackageTestRequest'] = _PACKAGETESTREQUEST
DESCRIPTOR.message_types_by_name['PackageTestResponse'] = _PACKAGETESTRESPONSE
DESCRIPTOR.message_types_by_name['PackageBuildRequest'] = _PACKAGEBUILDREQUEST
DESCRIPTOR.message_types_by_name['PackageBuildResponse'] = _PACKAGEBUILDRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

PackageListRequest = _reflection.GeneratedProtocolMessageType('PackageListRequest', (_message.Message,), {
  'DESCRIPTOR' : _PACKAGELISTREQUEST,
  '__module__' : 'PackageOperations_pb2'
  # @@protoc_insertion_point(class_scope:PackageOperations.PackageListRequest)
  })
_sym_db.RegisterMessage(PackageListRequest)

PackageList = _reflection.GeneratedProtocolMessageType('PackageList', (_message.Message,), {
  'DESCRIPTOR' : _PACKAGELIST,
  '__module__' : 'PackageOperations_pb2'
  # @@protoc_insertion_point(class_scope:PackageOperations.PackageList)
  })
_sym_db.RegisterMessage(PackageList)

PackageTestRequest = _reflection.GeneratedProtocolMessageType('PackageTestRequest', (_message.Message,), {
  'DESCRIPTOR' : _PACKAGETESTREQUEST,
  '__module__' : 'PackageOperations_pb2'
  # @@protoc_insertion_point(class_scope:PackageOperations.PackageTestRequest)
  })
_sym_db.RegisterMessage(PackageTestRequest)

PackageTestResponse = _reflection.GeneratedProtocolMessageType('PackageTestResponse', (_message.Message,), {
  'DESCRIPTOR' : _PACKAGETESTRESPONSE,
  '__module__' : 'PackageOperations_pb2'
  # @@protoc_insertion_point(class_scope:PackageOperations.PackageTestResponse)
  })
_sym_db.RegisterMessage(PackageTestResponse)

PackageBuildRequest = _reflection.GeneratedProtocolMessageType('PackageBuildRequest', (_message.Message,), {
  'DESCRIPTOR' : _PACKAGEBUILDREQUEST,
  '__module__' : 'PackageOperations_pb2'
  # @@protoc_insertion_point(class_scope:PackageOperations.PackageBuildRequest)
  })
_sym_db.RegisterMessage(PackageBuildRequest)

PackageBuildResponse = _reflection.GeneratedProtocolMessageType('PackageBuildResponse', (_message.Message,), {
  'DESCRIPTOR' : _PACKAGEBUILDRESPONSE,
  '__module__' : 'PackageOperations_pb2'
  # @@protoc_insertion_point(class_scope:PackageOperations.PackageBuildResponse)
  })
_sym_db.RegisterMessage(PackageBuildResponse)


DESCRIPTOR._options = None

_PACKAGEOPERATIONSERVICES = _descriptor.ServiceDescriptor(
  name='PackageOperationServices',
  full_name='PackageOperations.PackageOperationServices',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=268,
  serialized_end=572,
  methods=[
  _descriptor.MethodDescriptor(
    name='ListPackages',
    full_name='PackageOperations.PackageOperationServices.ListPackages',
    index=0,
    containing_service=None,
    input_type=_PACKAGELISTREQUEST,
    output_type=_PACKAGELIST,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='TestPackage',
    full_name='PackageOperations.PackageOperationServices.TestPackage',
    index=1,
    containing_service=None,
    input_type=_PACKAGETESTREQUEST,
    output_type=_PACKAGETESTRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='BuildPackage',
    full_name='PackageOperations.PackageOperationServices.BuildPackage',
    index=2,
    containing_service=None,
    input_type=_PACKAGEBUILDREQUEST,
    output_type=_PACKAGEBUILDRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_PACKAGEOPERATIONSERVICES)

DESCRIPTOR.services_by_name['PackageOperationServices'] = _PACKAGEOPERATIONSERVICES

# @@protoc_insertion_point(module_scope)
