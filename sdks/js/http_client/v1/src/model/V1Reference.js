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
 * OpenAPI spec version: 1.0.73
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
    define(['ApiClient', 'model/V1DagReference', 'model/V1HubReference', 'model/V1PathReference', 'model/V1UrlReference'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./V1DagReference'), require('./V1HubReference'), require('./V1PathReference'), require('./V1UrlReference'));
  } else {
    // Browser globals (root is window)
    if (!root.PolyaxonSdk) {
      root.PolyaxonSdk = {};
    }
    root.PolyaxonSdk.V1Reference = factory(root.PolyaxonSdk.ApiClient, root.PolyaxonSdk.V1DagReference, root.PolyaxonSdk.V1HubReference, root.PolyaxonSdk.V1PathReference, root.PolyaxonSdk.V1UrlReference);
  }
}(this, function(ApiClient, V1DagReference, V1HubReference, V1PathReference, V1UrlReference) {
  'use strict';

  /**
   * The V1Reference model module.
   * @module model/V1Reference
   * @version 1.0.73
   */

  /**
   * Constructs a new <code>V1Reference</code>.
   * @alias module:model/V1Reference
   * @class
   */
  var exports = function() {
  };

  /**
   * Constructs a <code>V1Reference</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/V1Reference} obj Optional instance to populate.
   * @return {module:model/V1Reference} The populated <code>V1Reference</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      if (data.hasOwnProperty('hub_reference'))
        obj.hub_reference = V1HubReference.constructFromObject(data['hub_reference']);
      if (data.hasOwnProperty('dag_reference'))
        obj.dag_reference = V1DagReference.constructFromObject(data['dag_reference']);
      if (data.hasOwnProperty('url_reference'))
        obj.url_reference = V1UrlReference.constructFromObject(data['url_reference']);
      if (data.hasOwnProperty('path_reference'))
        obj.path_reference = V1PathReference.constructFromObject(data['path_reference']);
    }
    return obj;
  }

  /**
   * @member {module:model/V1HubReference} hub_reference
   */
  exports.prototype.hub_reference = undefined;

  /**
   * @member {module:model/V1DagReference} dag_reference
   */
  exports.prototype.dag_reference = undefined;

  /**
   * @member {module:model/V1UrlReference} url_reference
   */
  exports.prototype.url_reference = undefined;

  /**
   * @member {module:model/V1PathReference} path_reference
   */
  exports.prototype.path_reference = undefined;

  return exports;

}));