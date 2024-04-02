
# How to reproduce

example for https://github.com/wundergraph/cosmo/issues/679

1.
```shell
make run
```

2.
```shell
npm install -g wgc@latest
wgc router compose -i graph.yaml -o config.json
```

3. use config.json for run router

4. run query

```graphql
query Test {
    unions {
        abUnion {
            __typename
            ... on A {
                ...AFragment
            }
            ... on B {
                ...BFragment
            }
        }
    }
}

fragment AFragment on A {
    __typename
    common {
        test
    }
}

fragment BFragment on B {
    __typename
    common {
        broken
    }
}
```

5. view result

```json
{
  "errors": [
    {
      "message": "Cannot return null for non-nullable field 'Query.unions.abUnion.common.broken'.",
      "path": [
        "unions",
        0,
        "abUnion",
        "common",
        "broken"
      ]
    }
  ],
  "data": null,
  "extensions": {
    "trace": {
      "info": {
        "trace_start_time": "2024-04-02T14:58:58+03:00",
        "trace_start_unix": 1712059138,
        "parse_stats": {
          "duration_nanoseconds": 388625,
          "duration_pretty": "388.625µs",
          "duration_since_start_nanoseconds": 157625,
          "duration_since_start_pretty": "157.625µs"
        },
        "normalize_stats": {
          "duration_nanoseconds": 739167,
          "duration_pretty": "739.167µs",
          "duration_since_start_nanoseconds": 565167,
          "duration_since_start_pretty": "565.167µs"
        },
        "validate_stats": {
          "duration_nanoseconds": 113292,
          "duration_pretty": "113.292µs",
          "duration_since_start_nanoseconds": 1314625,
          "duration_since_start_pretty": "1.314625ms"
        },
        "planner_stats": {
          "duration_nanoseconds": 4123166,
          "duration_pretty": "4.123166ms",
          "duration_since_start_nanoseconds": 1428584,
          "duration_since_start_pretty": "1.428584ms"
        }
      },
      "fetch": {
        "id": "335d69f3-9d4b-4d9d-8e1e-7b307be6c3bc",
        "type": "single",
        "path": "query",
        "data_source_id": "0",
        "datasource_load_trace": {
          "raw_input_data": {},
          "input": {
            "body": {
              "query": "{unions {abUnion {__typename ... on A {__typename common {test}} ... on B {__typename common {broken}}}}}"
            },
            "header": {},
            "method": "POST",
            "url": "http://localhost:8080/query"
          },
          "output": {
            "data": {
              "unions": [
                {
                  "abUnion": {
                    "__typename": "A",
                    "common": {
                      "test": 1
                    }
                  }
                },
                {
                  "abUnion": {
                    "__typename": "A",
                    "common": {
                      "test": 1
                    }
                  }
                }
              ]
            },
            "extensions": {
              "trace": {
                "request": {
                  "method": "POST",
                  "url": "http://localhost:8080/query",
                  "headers": {
                    "Accept": [
                      "application/json"
                    ],
                    "Accept-Encoding": [
                      "gzip",
                      "deflate"
                    ],
                    "Content-Type": [
                      "application/json"
                    ],
                    "Myheader": [
                      ""
                    ]
                  }
                },
                "response": {
                  "status_code": 200,
                  "status": "200 OK",
                  "headers": {
                    "Content-Length": [
                      "157"
                    ],
                    "Content-Type": [
                      "application/json"
                    ],
                    "Date": [
                      "Tue, 02 Apr 2024 11:58:58 GMT"
                    ]
                  },
                  "body_size": 157
                }
              }
            }
          },
          "duration_since_start_nanoseconds": 6022500,
          "duration_since_start_pretty": "6.0225ms",
          "duration_load_nanoseconds": 4066125,
          "duration_load_pretty": "4.066125ms",
          "single_flight_used": true,
          "single_flight_shared_response": false,
          "load_skipped": false,
          "load_stats": {
            "get_conn": {
              "duration_since_start_nanoseconds": 6902750,
              "duration_since_start_pretty": "6.90275ms",
              "host_port": "localhost:8080"
            },
            "got_conn": {
              "duration_since_start_nanoseconds": 6931834,
              "duration_since_start_pretty": "6.931834ms",
              "reused": true,
              "was_idle": true,
              "idle_time_nanoseconds": 811850625,
              "idle_time_pretty": "811.850625ms"
            },
            "got_first_response_byte": {
              "duration_since_start_nanoseconds": 7545042,
              "duration_since_start_pretty": "7.545042ms"
            },
            "dns_start": {
              "duration_since_start_nanoseconds": 0,
              "duration_since_start_pretty": "",
              "host": ""
            },
            "dns_done": {
              "duration_since_start_nanoseconds": 0,
              "duration_since_start_pretty": ""
            },
            "connect_start": {
              "duration_since_start_nanoseconds": 0,
              "duration_since_start_pretty": "",
              "network": "",
              "addr": ""
            },
            "connect_done": {
              "duration_since_start_nanoseconds": 0,
              "duration_since_start_pretty": "",
              "network": "",
              "addr": ""
            },
            "tls_handshake_start": {
              "duration_since_start_nanoseconds": 0,
              "duration_since_start_pretty": ""
            },
            "tls_handshake_done": {
              "duration_since_start_nanoseconds": 0,
              "duration_since_start_pretty": ""
            },
            "wrote_headers": {
              "duration_since_start_nanoseconds": 7048834,
              "duration_since_start_pretty": "7.048834ms"
            },
            "wrote_request": {
              "duration_since_start_nanoseconds": 7067667,
              "duration_since_start_pretty": "7.067667ms"
            }
          }
        }
      },
      "node_type": "object",
      "nullable": true,
      "fields": [
        {
          "name": "unions",
          "value": {
            "node_type": "array",
            "path": [
              "unions"
            ],
            "items": [
              {
                "node_type": "object",
                "nullable": true,
                "fields": [
                  {
                    "name": "abUnion",
                    "value": {
                      "node_type": "object",
                      "path": [
                        "abUnion"
                      ],
                      "fields": [
                        {
                          "name": "__typename",
                          "value": {
                            "node_type": "string",
                            "path": [
                              "__typename"
                            ],
                            "is_type_name": true
                          },
                          "parent_type_names": [
                            "ABUnion",
                            "A",
                            "B"
                          ],
                          "named_type": "String",
                          "data_source_ids": [
                            "0"
                          ]
                        },
                        {
                          "name": "common",
                          "value": {
                            "node_type": "object",
                            "path": [
                              "common"
                            ],
                            "fields": [
                              {
                                "name": "test",
                                "value": {
                                  "node_type": "integer",
                                  "path": [
                                    "test"
                                  ]
                                },
                                "parent_type_names": [
                                  "CommonA"
                                ],
                                "named_type": "Int",
                                "data_source_ids": [
                                  "0"
                                ]
                              },
                              {
                                "name": "broken",
                                "value": {
                                  "node_type": "integer",
                                  "path": [
                                    "broken"
                                  ]
                                },
                                "parent_type_names": [
                                  "CommonB"
                                ],
                                "named_type": "Int",
                                "data_source_ids": [
                                  "0"
                                ]
                              }
                            ]
                          },
                          "parent_type_names": [
                            "A",
                            "B"
                          ],
                          "named_type": "CommonA",
                          "data_source_ids": [
                            "0"
                          ]
                        }
                      ]
                    },
                    "parent_type_names": [
                      "Unions"
                    ],
                    "named_type": "ABUnion",
                    "data_source_ids": [
                      "0"
                    ]
                  }
                ]
              }
            ]
          },
          "parent_type_names": [
            "Query"
          ],
          "named_type": "Unions",
          "data_source_ids": [
            "0"
          ]
        }
      ]
    }
  }
}
```