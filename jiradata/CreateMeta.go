package jiradata

/////////////////////////////////////////////////////////////////////////
// This Code is Generated by SlipScheme Project:
// https://github.com/coryb/slipscheme
//
// Generated with command:
// slipscheme -dir jiradata -pkg jiradata -overwrite schemas/CreateMeta.json
/////////////////////////////////////////////////////////////////////////
//                            DO NOT EDIT                              //
/////////////////////////////////////////////////////////////////////////

// CreateMeta defined from schema:
// {
//   "title": "Create Meta",
//   "id": "https://docs.atlassian.com/jira/REST/schema/create-meta#",
//   "type": "object",
//   "properties": {
//     "expand": {
//       "title": "expand",
//       "type": "string"
//     },
//     "projects": {
//       "title": "projects",
//       "type": "array",
//       "items": {
//         "title": "Create Meta Project",
//         "type": "object",
//         "properties": {
//           "avatarUrls": {
//             "title": "avatarUrls",
//             "type": "object",
//             "patternProperties": {
//               ".+": {
//                 "type": "string"
//               }
//             }
//           },
//           "expand": {
//             "title": "expand",
//             "type": "string"
//           },
//           "id": {
//             "title": "id",
//             "type": "string"
//           },
//           "issuetypes": {
//             "title": "issuetypes",
//             "type": "array",
//             "items": {
//               "title": "Create Meta Issue Type",
//               "type": "object",
//               "properties": {
//                 "avatarId": {
//                   "title": "avatarId",
//                   "type": "integer"
//                 },
//                 "description": {
//                   "title": "description",
//                   "type": "string"
//                 },
//                 "expand": {
//                   "title": "expand",
//                   "type": "string"
//                 },
//                 "fields": {
//                   "title": "fields",
//                   "type": "object",
//                   "patternProperties": {
//                     ".+": {
//                       "title": "Field Meta",
//                       "type": "object",
//                       "properties": {
//                         "allowedValues": {
//                           "title": "allowedValues",
//                           "type": "array",
//                           "items": {}
//                         },
//                         "autoCompleteUrl": {
//                           "title": "autoCompleteUrl",
//                           "type": "string"
//                         },
//                         "defaultValue": {
//                           "title": "defaultValue"
//                         },
//                         "hasDefaultValue": {
//                           "title": "hasDefaultValue",
//                           "type": "boolean"
//                         },
//                         "key": {
//                           "title": "key",
//                           "type": "string"
//                         },
//                         "name": {
//                           "title": "name",
//                           "type": "string"
//                         },
//                         "operations": {
//                           "title": "operations",
//                           "type": "array",
//                           "items": {
//                             "type": "string"
//                           }
//                         },
//                         "required": {
//                           "title": "required",
//                           "type": "boolean"
//                         },
//                         "schema": {
//                           "title": "Json Type",
//                           "type": "object",
//                           "properties": {
//                             "custom": {
//                               "title": "custom",
//                               "type": "string"
//                             },
//                             "customId": {
//                               "title": "customId",
//                               "type": "integer"
//                             },
//                             "items": {
//                               "title": "items",
//                               "type": "string"
//                             },
//                             "system": {
//                               "title": "system",
//                               "type": "string"
//                             },
//                             "type": {
//                               "title": "type",
//                               "type": "string"
//                             }
//                           }
//                         }
//                       }
//                     }
//                   }
//                 },
//                 "iconUrl": {
//                   "title": "iconUrl",
//                   "type": "string"
//                 },
//                 "id": {
//                   "title": "id",
//                   "type": "string"
//                 },
//                 "name": {
//                   "title": "name",
//                   "type": "string"
//                 },
//                 "self": {
//                   "title": "self",
//                   "type": "string"
//                 },
//                 "subtask": {
//                   "title": "subtask",
//                   "type": "boolean"
//                 }
//               }
//             }
//           },
//           "key": {
//             "title": "key",
//             "type": "string"
//           },
//           "name": {
//             "title": "name",
//             "type": "string"
//           },
//           "self": {
//             "title": "self",
//             "type": "string"
//           }
//         }
//       }
//     }
//   }
// }
type CreateMeta struct {
	Expand   string   `json:"expand,omitempty" yaml:"expand,omitempty"`
	Projects Projects `json:"projects,omitempty" yaml:"projects,omitempty"`
}