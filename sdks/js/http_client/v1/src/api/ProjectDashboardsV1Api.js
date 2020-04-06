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
    define(['ApiClient', 'model/RuntimeError', 'model/V1Dashboard', 'model/V1ListDashboardsResponse'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('../model/RuntimeError'), require('../model/V1Dashboard'), require('../model/V1ListDashboardsResponse'));
  } else {
    // Browser globals (root is window)
    if (!root.PolyaxonSdk) {
      root.PolyaxonSdk = {};
    }
    root.PolyaxonSdk.ProjectDashboardsV1Api = factory(root.PolyaxonSdk.ApiClient, root.PolyaxonSdk.RuntimeError, root.PolyaxonSdk.V1Dashboard, root.PolyaxonSdk.V1ListDashboardsResponse);
  }
}(this, function(ApiClient, RuntimeError, V1Dashboard, V1ListDashboardsResponse) {
  'use strict';

  /**
   * ProjectDashboardsV1 service.
   * @module api/ProjectDashboardsV1Api
   * @version 1.0.71
   */

  /**
   * Constructs a new ProjectDashboardsV1Api. 
   * @alias module:api/ProjectDashboardsV1Api
   * @class
   * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
   * default to {@link module:ApiClient#instance} if unspecified.
   */
  var exports = function(apiClient) {
    this.apiClient = apiClient || ApiClient.instance;


    /**
     * Callback function to receive the result of the projectDashboardsV1CreateProjectDashboard operation.
     * @callback module:api/ProjectDashboardsV1Api~projectDashboardsV1CreateProjectDashboardCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Dashboard} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Create project dashboard
     * @param {String} owner Owner of the namespace
     * @param {String} project Project under namesapce
     * @param {module:model/V1Dashboard} body Dashboard body
     * @param {module:api/ProjectDashboardsV1Api~projectDashboardsV1CreateProjectDashboardCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Dashboard}
     */
    this.projectDashboardsV1CreateProjectDashboard = function(owner, project, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling projectDashboardsV1CreateProjectDashboard");
      }

      // verify the required parameter 'project' is set
      if (project === undefined || project === null) {
        throw new Error("Missing the required parameter 'project' when calling projectDashboardsV1CreateProjectDashboard");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling projectDashboardsV1CreateProjectDashboard");
      }


      var pathParams = {
        'owner': owner,
        'project': project
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
      var returnType = V1Dashboard;

      return this.apiClient.callApi(
        '/api/v1/{owner}/{project}/dashboards', 'POST',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the projectDashboardsV1DeleteProjectDashboard operation.
     * @callback module:api/ProjectDashboardsV1Api~projectDashboardsV1DeleteProjectDashboardCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete project dashboard
     * @param {String} owner Owner of the namespace
     * @param {String} project Project
     * @param {String} uuid Uuid identifier of the entity
     * @param {module:api/ProjectDashboardsV1Api~projectDashboardsV1DeleteProjectDashboardCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.projectDashboardsV1DeleteProjectDashboard = function(owner, project, uuid, callback) {
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling projectDashboardsV1DeleteProjectDashboard");
      }

      // verify the required parameter 'project' is set
      if (project === undefined || project === null) {
        throw new Error("Missing the required parameter 'project' when calling projectDashboardsV1DeleteProjectDashboard");
      }

      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling projectDashboardsV1DeleteProjectDashboard");
      }


      var pathParams = {
        'owner': owner,
        'project': project,
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
        '/api/v1/{owner}/{project}/dashboards/{uuid}', 'DELETE',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the projectDashboardsV1GetProjectDashboard operation.
     * @callback module:api/ProjectDashboardsV1Api~projectDashboardsV1GetProjectDashboardCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Dashboard} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get project dashboard
     * @param {String} owner Owner of the namespace
     * @param {String} project Project
     * @param {String} uuid Uuid identifier of the entity
     * @param {module:api/ProjectDashboardsV1Api~projectDashboardsV1GetProjectDashboardCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Dashboard}
     */
    this.projectDashboardsV1GetProjectDashboard = function(owner, project, uuid, callback) {
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling projectDashboardsV1GetProjectDashboard");
      }

      // verify the required parameter 'project' is set
      if (project === undefined || project === null) {
        throw new Error("Missing the required parameter 'project' when calling projectDashboardsV1GetProjectDashboard");
      }

      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling projectDashboardsV1GetProjectDashboard");
      }


      var pathParams = {
        'owner': owner,
        'project': project,
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
      var returnType = V1Dashboard;

      return this.apiClient.callApi(
        '/api/v1/{owner}/{project}/dashboards/{uuid}', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the projectDashboardsV1ListProjectDashboardNames operation.
     * @callback module:api/ProjectDashboardsV1Api~projectDashboardsV1ListProjectDashboardNamesCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ListDashboardsResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List project dashboard
     * @param {String} owner Owner of the namespace
     * @param {String} project Project under namesapce
     * @param {Object} opts Optional parameters
     * @param {Number} opts.offset Pagination offset.
     * @param {Number} opts.limit Limit size.
     * @param {String} opts.sort Sort to order the search.
     * @param {String} opts.query Query filter the search search.
     * @param {module:api/ProjectDashboardsV1Api~projectDashboardsV1ListProjectDashboardNamesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ListDashboardsResponse}
     */
    this.projectDashboardsV1ListProjectDashboardNames = function(owner, project, opts, callback) {
      opts = opts || {};
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling projectDashboardsV1ListProjectDashboardNames");
      }

      // verify the required parameter 'project' is set
      if (project === undefined || project === null) {
        throw new Error("Missing the required parameter 'project' when calling projectDashboardsV1ListProjectDashboardNames");
      }


      var pathParams = {
        'owner': owner,
        'project': project
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
      var returnType = V1ListDashboardsResponse;

      return this.apiClient.callApi(
        '/api/v1/{owner}/{project}/dashboards/names', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the projectDashboardsV1ListProjectDashboards operation.
     * @callback module:api/ProjectDashboardsV1Api~projectDashboardsV1ListProjectDashboardsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ListDashboardsResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List project dashboards
     * @param {String} owner Owner of the namespace
     * @param {String} project Project under namesapce
     * @param {Object} opts Optional parameters
     * @param {Number} opts.offset Pagination offset.
     * @param {Number} opts.limit Limit size.
     * @param {String} opts.sort Sort to order the search.
     * @param {String} opts.query Query filter the search search.
     * @param {module:api/ProjectDashboardsV1Api~projectDashboardsV1ListProjectDashboardsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ListDashboardsResponse}
     */
    this.projectDashboardsV1ListProjectDashboards = function(owner, project, opts, callback) {
      opts = opts || {};
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling projectDashboardsV1ListProjectDashboards");
      }

      // verify the required parameter 'project' is set
      if (project === undefined || project === null) {
        throw new Error("Missing the required parameter 'project' when calling projectDashboardsV1ListProjectDashboards");
      }


      var pathParams = {
        'owner': owner,
        'project': project
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
      var returnType = V1ListDashboardsResponse;

      return this.apiClient.callApi(
        '/api/v1/{owner}/{project}/dashboards', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the projectDashboardsV1PatchProjectDashboard operation.
     * @callback module:api/ProjectDashboardsV1Api~projectDashboardsV1PatchProjectDashboardCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Dashboard} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Patch project dashboard
     * @param {String} owner Owner of the namespace
     * @param {String} project Project under namesapce
     * @param {String} dashboard_uuid UUID
     * @param {module:model/V1Dashboard} body Dashboard body
     * @param {module:api/ProjectDashboardsV1Api~projectDashboardsV1PatchProjectDashboardCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Dashboard}
     */
    this.projectDashboardsV1PatchProjectDashboard = function(owner, project, dashboard_uuid, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling projectDashboardsV1PatchProjectDashboard");
      }

      // verify the required parameter 'project' is set
      if (project === undefined || project === null) {
        throw new Error("Missing the required parameter 'project' when calling projectDashboardsV1PatchProjectDashboard");
      }

      // verify the required parameter 'dashboard_uuid' is set
      if (dashboard_uuid === undefined || dashboard_uuid === null) {
        throw new Error("Missing the required parameter 'dashboard_uuid' when calling projectDashboardsV1PatchProjectDashboard");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling projectDashboardsV1PatchProjectDashboard");
      }


      var pathParams = {
        'owner': owner,
        'project': project,
        'dashboard.uuid': dashboard_uuid
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
      var returnType = V1Dashboard;

      return this.apiClient.callApi(
        '/api/v1/{owner}/{project}/dashboards/{dashboard.uuid}', 'PATCH',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the projectDashboardsV1PromoteProjectDashboard operation.
     * @callback module:api/ProjectDashboardsV1Api~projectDashboardsV1PromoteProjectDashboardCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Dashboard} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Promote project dashboard
     * @param {String} owner Owner of the namespace
     * @param {String} project Project under namesapce
     * @param {String} dashboard_uuid UUID
     * @param {module:api/ProjectDashboardsV1Api~projectDashboardsV1PromoteProjectDashboardCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Dashboard}
     */
    this.projectDashboardsV1PromoteProjectDashboard = function(owner, project, dashboard_uuid, callback) {
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling projectDashboardsV1PromoteProjectDashboard");
      }

      // verify the required parameter 'project' is set
      if (project === undefined || project === null) {
        throw new Error("Missing the required parameter 'project' when calling projectDashboardsV1PromoteProjectDashboard");
      }

      // verify the required parameter 'dashboard_uuid' is set
      if (dashboard_uuid === undefined || dashboard_uuid === null) {
        throw new Error("Missing the required parameter 'dashboard_uuid' when calling projectDashboardsV1PromoteProjectDashboard");
      }


      var pathParams = {
        'owner': owner,
        'project': project,
        'dashboard.uuid': dashboard_uuid
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
      var returnType = V1Dashboard;

      return this.apiClient.callApi(
        '/api/v1/{owner}/{project}/dashboards/{dashboard.uuid}/promote', 'POST',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the projectDashboardsV1UpdateProjectDashboard operation.
     * @callback module:api/ProjectDashboardsV1Api~projectDashboardsV1UpdateProjectDashboardCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Dashboard} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update project dashboard
     * @param {String} owner Owner of the namespace
     * @param {String} project Project under namesapce
     * @param {String} dashboard_uuid UUID
     * @param {module:model/V1Dashboard} body Dashboard body
     * @param {module:api/ProjectDashboardsV1Api~projectDashboardsV1UpdateProjectDashboardCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Dashboard}
     */
    this.projectDashboardsV1UpdateProjectDashboard = function(owner, project, dashboard_uuid, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling projectDashboardsV1UpdateProjectDashboard");
      }

      // verify the required parameter 'project' is set
      if (project === undefined || project === null) {
        throw new Error("Missing the required parameter 'project' when calling projectDashboardsV1UpdateProjectDashboard");
      }

      // verify the required parameter 'dashboard_uuid' is set
      if (dashboard_uuid === undefined || dashboard_uuid === null) {
        throw new Error("Missing the required parameter 'dashboard_uuid' when calling projectDashboardsV1UpdateProjectDashboard");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling projectDashboardsV1UpdateProjectDashboard");
      }


      var pathParams = {
        'owner': owner,
        'project': project,
        'dashboard.uuid': dashboard_uuid
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
      var returnType = V1Dashboard;

      return this.apiClient.callApi(
        '/api/v1/{owner}/{project}/dashboards/{dashboard.uuid}', 'PUT',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
  };

  return exports;
}));
