dst=dist

build:
	mkdir -p $(dst)/config

	gradle build
	cp app/build/libs/magnolia-app-*.jar $(dst)
	cp app/src/main/resources/application.properties $(dst)/config/
	cp app/src/main/resources/logback-file.xml $(dst)/config/logback.xml
	cp run.sh $(dst)


clean:
	gradle clean
	-rm -r $(dst)
