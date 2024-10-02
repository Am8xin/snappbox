Hi!
This is my answer to snappbox challenge which is a fare calculator the for multiple deliveries implemented in Golang.

My code have 3 folders(packages) for being clean and structural and a main.go file which acts as the main flow of the program. By the way It has 4 obvious parts included :

1. In main file, first I read them from a csv file and put all of them in an array of points. (using utils package)
2. Then I filter the dataset based on the conditions that It has been told for removing invalid points and also grouping the valid data by their delivery id.
3. After that , I calculate the fare of the each delivery based on the timestamp and haversine distance of each point of the delivery by calculating sum of the point to point fare.
4. And after all , I write the (id_delivery,fare_estimate) pair into a new csv file, if the fare_estimate is valid.

