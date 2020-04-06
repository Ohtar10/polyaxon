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
 * Do not edit the class manually.
 */


package io.swagger.client.model;

import java.util.Objects;
import java.util.Arrays;
import io.swagger.annotations.ApiModel;
import com.google.gson.annotations.SerializedName;

import java.io.IOException;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;

/**
 * &#x60;NullValue&#x60; is a singleton enumeration to represent the null value for the &#x60;Value&#x60; type union.   The JSON representation for &#x60;NullValue&#x60; is JSON &#x60;null&#x60;.   - NULL_VALUE: Null value.
 */
@JsonAdapter(ProtobufNullValue.Adapter.class)
public enum ProtobufNullValue {
  
  NULL_VALUE("NULL_VALUE");

  private String value;

  ProtobufNullValue(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static ProtobufNullValue fromValue(String text) {
    for (ProtobufNullValue b : ProtobufNullValue.values()) {
      if (String.valueOf(b.value).equals(text)) {
        return b;
      }
    }
    return null;
  }

  public static class Adapter extends TypeAdapter<ProtobufNullValue> {
    @Override
    public void write(final JsonWriter jsonWriter, final ProtobufNullValue enumeration) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public ProtobufNullValue read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return ProtobufNullValue.fromValue(String.valueOf(value));
    }
  }
}
