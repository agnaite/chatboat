/**
 * Copyright 2016 Confluent Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
#pragma once


/**
 * Glue between Go, Cgo and librdkafka
 */


/**
 * Represents a fetched C message, with all extra fields extracted
 * to struct fields.
 */
typedef struct fetched_c_msg {
  rd_kafka_message_t *msg;
  rd_kafka_timestamp_type_t tstype;
  int64_t ts;
} fetched_c_msg_t;
