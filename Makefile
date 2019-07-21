SERVICE_NAME=parking-lot-app
PROJECTS=$(SERVICE_NAME)/...

all:
	@cd src/ && go install $(PROJECTS)

test:
	@cd src/ && go test -coverprofile=coverage.out $(PROJECTS)
	@go tool cover -html=src/coverage.out

# generate XML report in Cobertura format
test.xml:
	@cd src/ && gocov test $(PROJECTS) | gocov-xml > coverage.xml



clean:
	@rm -rf ./bin/$(SERVICE_NAME)

clean-all:

run:
	@./bin/parking-lot-app
#	Run as backgroud process.
#	@./bin/parking-lot-app &

stop:
	# Needed, if the application is running as background process, 
	@pkill -f "./bin/parking-lot-app" &