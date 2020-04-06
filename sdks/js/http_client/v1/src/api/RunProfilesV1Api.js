// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * OpenAPI spec version: 1.0.71
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.10
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/RuntimeError', 'model/V1ListRunProfilesResponse', 'model/V1RunProfile'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('../model/RuntimeError'), require('../model/V1ListRunProfilesResponse'), require('../model/V1RunProfile'));
  } else {
    // Browser globals (root is window)
    if (!root.PolyaxonSdk) {
      root.PolyaxonSdk = {};
    }
    root.PolyaxonSdk.RunProfilesV1Api = factory(root.PolyaxonSdk.ApiClient, root.PolyaxonSdk.RuntimeError, root.PolyaxonSdk.V1ListRunProfilesResponse, root.PolyaxonSdk.V1RunProfile);
  }
}(this, function(ApiClient, RuntimeError, V1ListRunProfilesResponse, V1RunProfile) {
  'use strict';

  /**
   * RunProfilesV1 service.
   * @module api/RunProfilesV1Api
   * @version 1.0.71
   */

  /**
   * Constructs a new RunProfilesV1Api. 
   * @alias module:api/RunProfilesV1Api
   * @class
   * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
   * default to {@link module:ApiClient#instance} if unspecified.
   */
  var exports = function(apiClient) {
    this.apiClient = apiClient || ApiClient.instance;


    /**
     * Callback function to receive the result of the runProfilesV1CreateRunProfile operation.
     * @callback module:api/RunProfilesV1Api~runProfilesV1CreateRunProfileCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1RunProfile} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Create hub component
     * @param {String} owner Owner of the namespace
     * @param {module:model/V1RunProfile} body Artifact store body
     * @param {module:api/RunProfilesV1Api~runProfilesV1CreateRunProfileCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1RunProfile}
     */
    this.runProfilesV1CreateRunProfile = function(owner, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling runProfilesV1CreateRunProfile");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling runProfilesV1CreateRunProfile");
      }


      var pathParams = {
        'owner': owner
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1RunProfile;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/run_profiles', 'POST',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the runProfilesV1DeleteRunProfile operation.
     * @callback module:api/RunProfilesV1Api~runProfilesV1DeleteRunProfileCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete hub component
     * @param {String} owner Owner of the namespace
     * @param {String} uuid Uuid identifier of the entity
     * @param {module:api/RunProfilesV1Api~runProfilesV1DeleteRunProfileCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.runProfilesV1DeleteRunProfile = function(owner, uuid, callback) {
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling runProfilesV1DeleteRunProfile");
      }

      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling runProfilesV1DeleteRunProfile");
      }


      var pathParams = {
        'owner': owner,
        'uuid': uuid
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = null;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/run_profiles/{uuid}', 'DELETE',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the runProfilesV1GetRunProfile operation.
     * @callback module:api/RunProfilesV1Api~runProfilesV1GetRunProfileCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1RunProfile} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get hub component
     * @param {String} owner Owner of the namespace
     * @param {String} uuid Uuid identifier of the entity
     * @param {module:api/RunProfilesV1Api~runProfilesV1GetRunProfileCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1RunProfile}
     */
    this.runProfilesV1GetRunProfile = function(owner, uuid, callback) {
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling runProfilesV1GetRunProfile");
      }

      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling runProfilesV1GetRunProfile");
      }


      var pathParams = {
        'owner': owner,
        'uuid': uuid
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1RunProfile;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/run_profiles/{uuid}', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the runProfilesV1ListRunProfileNames operation.
     * @callback module:api/RunProfilesV1Api~runProfilesV1ListRunProfileNamesCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ListRunProfilesResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List hub component names
     * @param {String} owner Owner of the namespace
     * @param {Object} opts Optional parameters
     * @param {Number} opts.offset Pagination offset.
     * @param {Number} opts.limit Limit size.
     * @param {String} opts.sort Sort to order the search.
     * @param {String} opts.query Query filter the search search.
     * @param {module:api/RunProfilesV1Api~runProfilesV1ListRunProfileNamesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ListRunProfilesResponse}
     */
    this.runProfilesV1ListRunProfileNames = function(owner, opts, callback) {
      opts = opts || {};
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling runProfilesV1ListRunProfileNames");
      }


      var pathParams = {
        'owner': owner
      };
      var queryParams = {
        'offset': opts['offset'],
        'limit': opts['limit'],
        'sort': opts['sort'],
        'query': opts['query'],
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1ListRunProfilesResponse;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/run_profiles/names', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the runProfilesV1ListRunProfiles operation.
     * @callback module:api/RunProfilesV1Api~runProfilesV1ListRunProfilesCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ListRunProfilesResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List hub components
     * @param {String} owner Owner of the namespace
     * @param {Object} opts Optional parameters
     * @param {Number} opts.offset Pagination offset.
     * @param {Number} opts.limit Limit size.
     * @param {String} opts.sort Sort to order the search.
     * @param {String} opts.query Query filter the search search.
     * @param {module:api/RunProfilesV1Api~runProfilesV1ListRunProfilesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ListRunProfilesResponse}
     */
    this.runProfilesV1ListRunProfiles = function(owner, opts, callback) {
      opts = opts || {};
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling runProfilesV1ListRunProfiles");
      }


      var pathParams = {
        'owner': owner
      };
      var queryParams = {
        'offset': opts['offset'],
        'limit': opts['limit'],
        'sort': opts['sort'],
        'query': opts['query'],
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1ListRunProfilesResponse;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/run_profiles', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the runProfilesV1PatchRunProfile operation.
     * @callback module:api/RunProfilesV1Api~runProfilesV1PatchRunProfileCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1RunProfile} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Patch hub component
     * @param {String} owner Owner of the namespace
     * @param {String} run_profile_uuid UUID
     * @param {module:model/V1RunProfile} body Artifact store body
     * @param {module:api/RunProfilesV1Api~runProfilesV1PatchRunProfileCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1RunProfile}
     */
    this.runProfilesV1PatchRunProfile = function(owner, run_profile_uuid, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling runProfilesV1PatchRunProfile");
      }

      // verify the required parameter 'run_profile_uuid' is set
      if (run_profile_uuid === undefined || run_profile_uuid === null) {
        throw new Error("Missing the required parameter 'run_profile_uuid' when calling runProfilesV1PatchRunProfile");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling runProfilesV1PatchRunProfile");
      }


      var pathParams = {
        'owner': owner,
        'run_profile.uuid': run_profile_uuid
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1RunProfile;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/run_profiles/{run_profile.uuid}', 'PATCH',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the runProfilesV1UpdateRunProfile operation.
     * @callback module:api/RunProfilesV1Api~runProfilesV1UpdateRunProfileCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1RunProfile} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update hub component
     * @param {String} owner Owner of the namespace
     * @param {String} run_profile_uuid UUID
     * @param {module:model/V1RunProfile} body Artifact store body
     * @param {module:api/RunProfilesV1Api~runProfilesV1UpdateRunProfileCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1RunProfile}
     */
    this.runProfilesV1UpdateRunProfile = function(owner, run_profile_uuid, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling runProfilesV1UpdateRunProfile");
      }

      // verify the required parameter 'run_profile_uuid' is set
      if (run_profile_uuid === undefined || run_profile_uuid === null) {
        throw new Error("Missing the required parameter 'run_profile_uuid' when calling runProfilesV1UpdateRunProfile");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling runProfilesV1UpdateRunProfile");
      }


      var pathParams = {
        'owner': owner,
        'run_profile.uuid': run_profile_uuid
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1RunProfile;

      return this.apiClient.callApi(
        '/api/v1/orgs/{owner}/run_profiles/{run_profile.uuid}', 'PUT',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
  };

  return exports;
}));
