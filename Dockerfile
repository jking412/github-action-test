FROM fabric8/java-alpine-openjdk11-jre

COPY target/*.jar /deployments/app.jar

EXPOSE 8000

ENTRYPOINT ["java", "-jar", "deployments/app.jar","--server.port=8000"]