### USAGE:

There are 3 type of code generation:
1. GraphQL Schema Generation
2. API code generation for Service. 
3. GRPC Client API code generation for API-Gateway

##### Methods:
`GenerateSchema()` for schema generation
`GenerateServiceAPI()` for service API generation
`GenerateClientGrpcAPI()` for Client GRPC API Code generation

Note: `LoadTemplates()` will load the templates in working project generation directory. This is required, since code generation is running from local project, it require
the templates to be in its local scope, this will copy the templates defined in this directory in local project directory in speficied path. 

#### 
