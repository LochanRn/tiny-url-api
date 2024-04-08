MAKEFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_DIR := $(dir $(MAKEFILE_PATH))

specs := ${CURRENT_DIR}api/spec/info.yaml ${CURRENT_DIR}api/spec/flatten_paths.yaml ${CURRENT_DIR}api/spec/system_v1.yaml

flatten:
	swagger30 flatten ${CURRENT_DIR}api/spec/paths.yaml --format=yaml -o ${CURRENT_DIR}api/spec/flatten_paths.yaml

version:
	@echo "############################################ *** SWAGGER VERSION *** ######################################################"
	swagger30 version
	@echo "###########################################################################################################################"

mixin: flatten version
	echo ${specs}
	swagger30 mixin ${specs} --format yaml -o ${CURRENT_DIR}api/spec/tiny-url-server_api.yaml
	rm ${CURRENT_DIR}api/spec/flatten_paths.yaml

generate: swagger-clean mixin
	swagger30 generate client --target=${CURRENT_DIR}/gen --spec=${CURRENT_DIR}/api/spec/tiny-url-server_api.yaml --name=tiny-url-server

swagger-clean:
	echo "Swagger Files and Folders will be cleaned"
	@[ -d ${CURRENT_DIR}gen/models ] && rm -rf ${CURRENT_DIR}gen/models/ || echo "Ignoring: Directory ${CURRENT_DIR}gen/models  does not exists."
	@[ -d ${CURRENT_DIR}gen/restapi/operations ] && rm -rf ${CURRENT_DIR}gen/restapi/operations/ || echo "Ignoring: Directory ${CURRENT_DIR}gen/restapi/operations  does not exists."
	@[ -f ${CURRENT_DIR}gen/restapi/doc.go ] && rm -rf ${CURRENT_DIR}gen/restapi/doc.go || echo "Ignoring: File ${CURRENT_DIR}gen/restapi/doc.go  does not exists."
	@[ -f ${CURRENT_DIR}gen/restapi/embedded_spec.go ] && rm -rf ${CURRENT_DIR}gen/restapi/embedded_spec.go || echo "Ignoring: File ${CURRENT_DIR}gen/restapi/embedded_spec.go  does not exists."
	@[ -f ${CURRENT_DIR}gen/restapi/server.go ] && rm -rf ${CURRENT_DIR}gen/restapi/server.go || echo "Ignoring: File ${CURRENT_DIR}gen/restapi/server.go  does not exists."
	echo "Swagger Files and Folders are cleaned"