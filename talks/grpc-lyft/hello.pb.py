// SERVER OMIT
@blueprint.route('/api/gethello', methods=['GET'])
def get_hello():
    if request.headers.get('Content-Type') == 'application/proto': // HL
        try:
            input = SayHelloRequest()
            input.ParseFromString(request.data) // HL

            # Call the actual implementation method
            resp = handle_hello_world_get(input) // HL
            return resp.SerializeToString() // HL
        except Exception as e:
            logger.warning(
                'Exception calling handle_hello_world_get on get_hello: {}'.format(repr(e))
            )
            raise e
    else:
        # Non proto application code goes here
        return handle_hello_world_get(request) // HL
// END SERVER OMIT

// CLIENT OMIT
def get_hello(self, input):
    try:
        assert isinstance(input, SayHelloRequest)
        headers = {
            'Content-Type': 'application/proto' // HL
        }
        response = self.get(
            '/api/gethello', // HL
            data=input.SerializeToString(), // HL
            headers=headers, // HL
            raw_request=True, // HL
            raw_response=True) // HL
        op = SayHelloResponse()
        op.ParseFromString(response.content) // HL

        return op
    except Exception as e:
        logger.warning(
            'Exception calling get_hello : {}'.format(repr(e))
        )
        raise e
// END CLIENT OMIT
