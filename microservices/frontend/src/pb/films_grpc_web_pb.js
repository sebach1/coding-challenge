/**
 * @fileoverview gRPC-Web generated client stub for films
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')
const proto = {};
proto.films = require('./films_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.films.FilmsClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.films.FilmsPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.films.RetrieveFilmRequest,
 *   !proto.films.RetrieveFilmResponse>}
 */
const methodDescriptor_Films_RetrieveFilm = new grpc.web.MethodDescriptor(
  '/films.Films/RetrieveFilm',
  grpc.web.MethodType.UNARY,
  proto.films.RetrieveFilmRequest,
  proto.films.RetrieveFilmResponse,
  /**
   * @param {!proto.films.RetrieveFilmRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.films.RetrieveFilmResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.films.RetrieveFilmRequest,
 *   !proto.films.RetrieveFilmResponse>}
 */
const methodInfo_Films_RetrieveFilm = new grpc.web.AbstractClientBase.MethodInfo(
  proto.films.RetrieveFilmResponse,
  /**
   * @param {!proto.films.RetrieveFilmRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.films.RetrieveFilmResponse.deserializeBinary
);


/**
 * @param {!proto.films.RetrieveFilmRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.films.RetrieveFilmResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.films.RetrieveFilmResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.films.FilmsClient.prototype.retrieveFilm =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/films.Films/RetrieveFilm',
      request,
      metadata || {},
      methodDescriptor_Films_RetrieveFilm,
      callback);
};


/**
 * @param {!proto.films.RetrieveFilmRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.films.RetrieveFilmResponse>}
 *     Promise that resolves to the response
 */
proto.films.FilmsPromiseClient.prototype.retrieveFilm =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/films.Films/RetrieveFilm',
      request,
      metadata || {},
      methodDescriptor_Films_RetrieveFilm);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.films.RetrieveFilmsRequest,
 *   !proto.films.RetrieveFilmsResponse>}
 */
const methodDescriptor_Films_RetrieveFilms = new grpc.web.MethodDescriptor(
  '/films.Films/RetrieveFilms',
  grpc.web.MethodType.UNARY,
  proto.films.RetrieveFilmsRequest,
  proto.films.RetrieveFilmsResponse,
  /**
   * @param {!proto.films.RetrieveFilmsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.films.RetrieveFilmsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.films.RetrieveFilmsRequest,
 *   !proto.films.RetrieveFilmsResponse>}
 */
const methodInfo_Films_RetrieveFilms = new grpc.web.AbstractClientBase.MethodInfo(
  proto.films.RetrieveFilmsResponse,
  /**
   * @param {!proto.films.RetrieveFilmsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.films.RetrieveFilmsResponse.deserializeBinary
);


/**
 * @param {!proto.films.RetrieveFilmsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.films.RetrieveFilmsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.films.RetrieveFilmsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.films.FilmsClient.prototype.retrieveFilms =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/films.Films/RetrieveFilms',
      request,
      metadata || {},
      methodDescriptor_Films_RetrieveFilms,
      callback);
};


/**
 * @param {!proto.films.RetrieveFilmsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.films.RetrieveFilmsResponse>}
 *     Promise that resolves to the response
 */
proto.films.FilmsPromiseClient.prototype.retrieveFilms =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/films.Films/RetrieveFilms',
      request,
      metadata || {},
      methodDescriptor_Films_RetrieveFilms);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.films.RetrieveFilmsWithPeopleRequest,
 *   !proto.films.RetrieveFilmsWithPeopleResponse>}
 */
const methodDescriptor_Films_RetrieveFilmsWithPeople = new grpc.web.MethodDescriptor(
  '/films.Films/RetrieveFilmsWithPeople',
  grpc.web.MethodType.UNARY,
  proto.films.RetrieveFilmsWithPeopleRequest,
  proto.films.RetrieveFilmsWithPeopleResponse,
  /**
   * @param {!proto.films.RetrieveFilmsWithPeopleRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.films.RetrieveFilmsWithPeopleResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.films.RetrieveFilmsWithPeopleRequest,
 *   !proto.films.RetrieveFilmsWithPeopleResponse>}
 */
const methodInfo_Films_RetrieveFilmsWithPeople = new grpc.web.AbstractClientBase.MethodInfo(
  proto.films.RetrieveFilmsWithPeopleResponse,
  /**
   * @param {!proto.films.RetrieveFilmsWithPeopleRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.films.RetrieveFilmsWithPeopleResponse.deserializeBinary
);


/**
 * @param {!proto.films.RetrieveFilmsWithPeopleRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.films.RetrieveFilmsWithPeopleResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.films.RetrieveFilmsWithPeopleResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.films.FilmsClient.prototype.retrieveFilmsWithPeople =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/films.Films/RetrieveFilmsWithPeople',
      request,
      metadata || {},
      methodDescriptor_Films_RetrieveFilmsWithPeople,
      callback);
};


/**
 * @param {!proto.films.RetrieveFilmsWithPeopleRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.films.RetrieveFilmsWithPeopleResponse>}
 *     Promise that resolves to the response
 */
proto.films.FilmsPromiseClient.prototype.retrieveFilmsWithPeople =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/films.Films/RetrieveFilmsWithPeople',
      request,
      metadata || {},
      methodDescriptor_Films_RetrieveFilmsWithPeople);
};


module.exports = proto.films;

