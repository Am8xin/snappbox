Hi! I am Mohammadamin Mahmoudi and this is a README file for the snappbox challenge and the documentatoin is in another .pdf file in the same folder ,So Let's Go :)

This is my answer to snappbox challenge which is a fare calculator the for multiple deliveries implemented in Golang using concurrency features of the Golang.

My project has 2 folders(packages) for being clean and structural and a main.go file which controls the main flow of the program. By the way It has 5 obvious parts included :

1. First of all , we should have a point data structure to store our data in a clean way. (models package)
2. In main file, first I read the input data from a csv file and put all of them in a map data structure which each delivery id is mapped to an array of points . (using utils package)
3. Then I filter the dataset based on the conditions that It has been told for removing invalid points and also grouping the valid data by their delivery id.
4. After that , I calculate the fare of the each delivery based on the timestamp and haversine distance of each point of the delivery by calculating sum of the point to point fare.
5. And after all , I write the (id_delivery,fare_estimate) pair into a new csv file, if the fare_estimate is valid.

I illustrate my code very explicit in a technical manner in the doc.pdf file which is in the same folder 