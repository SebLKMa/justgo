Recipe Stats Calculator
====

This project aims to process an automatically generated JSON file with recipe data and calculated some stats.

Instructions for collaborators
-----

1. Clone this repository.
2. Create a new branch called `dev`.
3. Create a pull request from your `dev` branch to the master branch.
4. Reply to the thread you're having with our HR department telling them we can start reviewing your code

Given
-----

Json fixtures file with recipe data. Download [Link](https://drive.google.com/file/d/19xg_U8R00RuVPuzED6I7xeFPWEiBQaiB/view?usp=sharing)

_Important notes_

1. Property value `"delivery"` always has the following format: "{weekday} {h}AM - {h}PM", i.e. "Monday 9AM - 5PM"
2. The number of distinct postcodes is lower than `1M`, one postcode is not longer than `10` chars.
3. The number of distinct recipe names is lower than `2K`, one recipe name is not longer than `100` chars.

Functional Requirements
------

1. Count the number of unique recipe names.
2. Count the number of occurences for each unique recipe name (alphabetically ordered by recipe name).
3. Find the postcode with most delivered recipes.
4. Count the number of deliveries to postcode `10120` that lie within the delivery time between `10AM` and `3PM`, examples _(`12AM` denotes midnight)_:
    - `NO` - `9AM - 2PM`
    - `YES` - `10AM - 2PM`
5. List the recipe names (alphabetically ordered) that contain in their name one of the following words:
    - Potato
    - Veggie
    - Mushroom

Non-functional Requirements
--------

1. The application is packaged with [Docker](https://www.docker.com/).
2. Setup scripts are provided.
3. The submission is provided as a `CLI` application.
4. The expected output is rendered to `stdout`. Make sure to render only the final `json`. If you need to print additional info or debug, pipe it to `stderr`.
5. It should be possible to (implementation is up to you):  
    a. provide a custom fixtures file as input  
    b. provide custom recipe names to search by (functional reqs. 5)  
    c. provide custom postcode and time window for search (functional reqs. 4)  

Expected output
---------------

Generate a JSON file of the following format:

```json5
{
    "unique_recipe_count": 15,
    "count_per_recipe": [
        {
            "recipe": "Mediterranean Baked Veggies",
            "count": 1
        },
        {
            "recipe": "Speedy Steak Fajitas",
            "count": 1
        },
        {
            "recipe": "Tex-Mex Tilapia",
            "count": 3
        }
    ],
    "busiest_postcode": {
        "postcode": "10120",
        "delivery_count": 1000
    },
    "count_per_postcode_and_time": {
        "postcode": "10120",
        "from": "11AM",
        "to": "3PM",
        "delivery_count": 500
    },
    "match_by_name": [
        "Mediterranean Baked Veggies", "Speedy Steak Fajitas", "Tex-Mex Tilapia"
    ]
}
```

Unit Tests Criteria
---

Test your submission against different input data sets - valid and invalid data sets.  

__General criteria from most important to less important__:

1. Functional and non-functional requirements are met.
2. Prefer application efficiency over code organisation complexity.
3. Code is readable and comprehensible. Setup instructions and run instructions are provided.
4. Tests are showcased (_no need to cover everything_).
5. Supporting notes on taken decisions and further clarifications are welcome.


Author's Comment
---

## Pre-requisite

Ubuntu Linux 18.04  
GO programming language development environment  
Docker  

### Checked-in files  
File | Description
------------ | -------------
hello.go | The main program 
hello_test.go | Some sample unit tests (just run `go test`)  
Dockerfile.alpine | Dockerfile for building the go binary for alpine  
buildalpine.sh | Builds Dockerfile.alpine  
dockerrun.sh | Just a simple script to demo run hello binary. Actual production may use docker mount volumes as input and output storage.

### Non-checked-in files  
File | Description
------------ | -------------
hello | Binary to be built simply using `go build` or by buildalpine.sh for Docker  
input.json | The input file for hello program. untar from [Link](https://test-golang-recipes.s3-eu-west-1.amazonaws.com/recipe-calculation-test-fixtures/hf_test_calculation_fixtures.tar.gz)

## Run Instructions

You can use `go run` during development and test, or build the binary.  
For consistency with deployment, we use the built binary.  

### Build binary

Build the go binary.  
```sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hello hello.go
```

### Run binary

Untar the file from [Link](https://test-golang-recipes.s3-eu-west-1.amazonaws.com/recipe-calculation-test-fixtures/hf_test_calculation_fixtures.tar.gz) as **input.json**  

Run the built go binary.  

```sh
./hello -input_file=input.json -output_file=output.json -query_postcode=10220 -query_FromHr=10AM -query_ToHr=6PM  -word_list=Chicken -word_list=Pork -word_list=Pasta
```

Example:  
```sh
./hello -input_file=input.json -output_file=output.json -query_postcode=10220 -query_FromHr=10AM -query_ToHr=6PM  -word_list=Chicken -word_list=Pork -word_list=Pasta

Your arguments -
  inputFileArg : input.json
  outputFileArg: output.json
  postcodeArg  : 10220
  fromHrArg    : 10AM
  toHrArg      : 6PM
  words        : [Chicken Pork Pasta]
2020/09/09 19:13:12 Starting Main loop
2020/09/09 19:13:28 
2020/09/09 19:13:28 Parsing input file took [15.480254356s]
Total of [10000000] object created.
2020/09/09 19:13:28 1. Number of recipes [29]
```

### Check the output

**Insights**: Input file could be large in size, please observe that the program's **memory consumption** is consistently conservative.  

The above command outputs to output.json file.  

Sample output:  
```sh
cat output.json 
{
  "unique_recipe_count": 29,
  "count_per_recipe": [
    {
      "recipe": "Cajun-Spiced Pulled Pork",
      "count": 667365
    },
    {
      "recipe": "Cheesy Chicken Enchilada Bake",
      "count": 333012
    },
    {
      "recipe": "Cherry Balsamic Pork Chops",
      "count": 333889
    },
    {
      "recipe": "Chicken Pineapple Quesadillas",
      "count": 331514
    },
    {
      "recipe": "Chicken Sausage Pizzas",
      "count": 333306
    },
    {
      "recipe": "Creamy Dill Chicken",
      "count": 333103
    },
    {
      "recipe": "Creamy Shrimp Tagliatelle",
      "count": 333395
    },
    {
      "recipe": "Crispy Cheddar Frico Cheeseburgers",
      "count": 333251
    },
    {
      "recipe": "Garden Quesadillas",
      "count": 333284
    },
    {
      "recipe": "Garlic Herb Butter Steak",
      "count": 333649
    },
    {
      "recipe": "Grilled Cheese and Veggie Jumble",
      "count": 333742
    },
    {
      "recipe": "Hearty Pork Chili",
      "count": 333355
    },
    {
      "recipe": "Honey Sesame Chicken",
      "count": 333748
    },
    {
      "recipe": "Hot Honey Barbecue Chicken Legs",
      "count": 334409
    },
    {
      "recipe": "Korean-Style Chicken Thighs",
      "count": 333069
    },
    {
      "recipe": "Meatloaf à La Mom",
      "count": 333570
    },
    {
      "recipe": "Mediterranean Baked Veggies",
      "count": 332939
    },
    {
      "recipe": "Melty Monterey Jack Burgers",
      "count": 333264
    },
    {
      "recipe": "Mole-Spiced Beef Tacos",
      "count": 332993
    },
    {
      "recipe": "One-Pan Orzo Italiano",
      "count": 333109
    },
    {
      "recipe": "Parmesan-Crusted Pork Tenderloin",
      "count": 333311
    },
    {
      "recipe": "Spanish One-Pan Chicken",
      "count": 333291
    },
    {
      "recipe": "Speedy Steak Fajitas",
      "count": 333578
    },
    {
      "recipe": "Spinach Artichoke Pasta Bake",
      "count": 333545
    },
    {
      "recipe": "Steakhouse-Style New York Strip",
      "count": 333473
    },
    {
      "recipe": "Stovetop Mac 'N' Cheese",
      "count": 333098
    },
    {
      "recipe": "Sweet Apple Pork Tenderloin",
      "count": 332595
    },
    {
      "recipe": "Tex-Mex Tilapia",
      "count": 333749
    },
    {
      "recipe": "Yellow Squash Flatbreads",
      "count": 333394
    }
  ],
  "busiest_postcode": {
    "postcode": "10176",
    "delivery_count": 91785
  },
  "count_per_postcode_and_time": {
    "postcode": "10220",
    "from": "10AM",
    "to": "6PM",
    "delivery_count": 44708
  },
  "match_by_name": [
    "Cajun-Spiced Pulled Pork",
    "Cheesy Chicken Enchilada Bake",
    "Cherry Balsamic Pork Chops",
    "Chicken Pineapple Quesadillas",
    "Chicken Sausage Pizzas",
    "Creamy Dill Chicken",
    "Hearty Pork Chili",
    "Honey Sesame Chicken",
    "Hot Honey Barbecue Chicken Legs",
    "Korean-Style Chicken Thighs",
    "Parmesan-Crusted Pork Tenderloin",
    "Spanish One-Pan Chicken",
    "Spinach Artichoke Pasta Bake",
    "Sweet Apple Pork Tenderloin"
  ]
}
```

### More sample command line argumemts

Use -h flag to list the command line arguments.  
```sh
./hello -h
Usage of ./hello:
  -input_file string
        The input JSON file to be process, defaults to input.json (default "input.json")
  -output_file string
        The output JSON file, defaults to output.json (default "output.json")
  -query_FromHr string
        postcode to query from hour, defaults to 9AM (default "9AM")
  -query_ToHr string
        postcode to query from hour, defaults to 11PM (default "11PM")
  -query_postcode string
        postcode to query, defaults to 10161 (default "10161")
  -word_list value
        query recipe containing these words
```

Sample uages:  
```sh
./hello --word_list=Chicken -word_list=Pork
./hello -query_postcode=10220 -query_FromHr=10AM -query_ToHr=6PM  -word_list=Chicken -word_list=Pork
./hello -query_postcode=10220 -query_FromHr=10AM -query_ToHr=6PM  -word_list=Chicken -word_list=Pork -word_list=Pasta
./hello -input_file=input.json -output_file=output.json -query_postcode=10220 -query_FromHr=10AM -query_ToHr=6PM  -word_list=Chicken -word_list=Pork -word_list=Pasta
```

## Docker container 

### Build Instructions

Run `buildalpine.sh` to build the go binary and docker image.  
```sh
./buildalpine.sh 

...
Remove one or more images
Untagged: hellofresh-hello-alpine:latest
Deleted: sha256:b5c1ce713069e3771f7b83a94617bd387b0dfdc18755fd947ca9eedc9f540952
Deleted: sha256:fd7ae522fba9b9f02753c09789ec7aee2ee3254eee474cb927468a365ffce510
Deleted: sha256:22d27d89398c1f5b2446eb2f8bfdd7c6f905c7e6d157ae2535f634e1d27985f4
Sending build context to Docker daemon  1.037GB
Step 1/5 : FROM alpine:latest
 ---> f70734b6a266
Step 2/5 : RUN mkdir /app
 ---> Using cache
 ---> b69242ef6f2a
Step 3/5 : WORKDIR /app
 ---> Using cache
 ---> 74a1e60d3f51
Step 4/5 : COPY hello .
 ---> de7097e571bd
Step 5/5 : CMD ["/bin/sh"]
 ---> Running in abc396b65762
Removing intermediate container abc396b65762
 ---> b8a3bf5cf364
Successfully built b8a3bf5cf364
Successfully tagged hellofresh-hello-alpine:latest
```

Verify docker image.  
```sh
docker images hellofresh-hello-alpine

REPOSITORY                TAG                 IMAGE ID            CREATED             SIZE
hellofresh-hello-alpine   latest              b8a3bf5cf364        43 seconds ago      8.68MB
```

### Run docker container

This is only a demo implementation demonstrating running Docker container to process input and output file.  
The actual implementation should use [docker mount or docker volume](https://docs.docker.com/storage/volumes/) as input and output storage.

For this demo, run `./dockerrun.sh`  
The demo is just using the default command line arguments.
```sh
./dockerrun.sh 
removing exited containers...
"docker rm" requires at least 1 argument.
See 'docker rm --help'.

Usage:  docker rm [OPTIONS] CONTAINER [CONTAINER...]

Remove one or more containers
start docker app and stay in background...
264aaa3f35c6eb828f16d631e1f2b11bb2ca0a2f19ebbd79c82dc232d3e0e02c
start docker app up and running
copying input file to container...
executing program...
Your arguments -
  inputFileArg: input.json
  outputFileArg: output.json
  postcodeArg: 10161
  fromHrArg: 9AM
  toHrArg: 11PM
  words: []
2020/09/10 03:02:34 Starting Main loop
2020/09/10 03:02:50 
2020/09/10 03:02:50 Parsing input file took [16.575122677s]
Total of [10000000] object created.
2020/09/10 03:02:50 1. Number of recipes [29]
copying output file from container...
done!
stopping docker app...
hellofresh
docker app stopped
```

### Check the output

Note that `"match_by_name": null` because we have not set any default words to search for recipes.  
Further enhancements could be made to pass arguments into Docker run.:blush: :technologist:  

```sh
cat output.json 
{
  "unique_recipe_count": 29,
  "count_per_recipe": [
    {
      "recipe": "Cajun-Spiced Pulled Pork",
      "count": 667365
    },
    {
      "recipe": "Cheesy Chicken Enchilada Bake",
      "count": 333012
    },
    {
      "recipe": "Cherry Balsamic Pork Chops",
      "count": 333889
    },
    {
      "recipe": "Chicken Pineapple Quesadillas",
      "count": 331514
    },
    {
      "recipe": "Chicken Sausage Pizzas",
      "count": 333306
    },
    {
      "recipe": "Creamy Dill Chicken",
      "count": 333103
    },
    {
      "recipe": "Creamy Shrimp Tagliatelle",
      "count": 333395
    },
    {
      "recipe": "Crispy Cheddar Frico Cheeseburgers",
      "count": 333251
    },
    {
      "recipe": "Garden Quesadillas",
      "count": 333284
    },
    {
      "recipe": "Garlic Herb Butter Steak",
      "count": 333649
    },
    {
      "recipe": "Grilled Cheese and Veggie Jumble",
      "count": 333742
    },
    {
      "recipe": "Hearty Pork Chili",
      "count": 333355
    },
    {
      "recipe": "Honey Sesame Chicken",
      "count": 333748
    },
    {
      "recipe": "Hot Honey Barbecue Chicken Legs",
      "count": 334409
    },
    {
      "recipe": "Korean-Style Chicken Thighs",
      "count": 333069
    },
    {
      "recipe": "Meatloaf à La Mom",
      "count": 333570
    },
    {
      "recipe": "Mediterranean Baked Veggies",
      "count": 332939
    },
    {
      "recipe": "Melty Monterey Jack Burgers",
      "count": 333264
    },
    {
      "recipe": "Mole-Spiced Beef Tacos",
      "count": 332993
    },
    {
      "recipe": "One-Pan Orzo Italiano",
      "count": 333109
    },
    {
      "recipe": "Parmesan-Crusted Pork Tenderloin",
      "count": 333311
    },
    {
      "recipe": "Spanish One-Pan Chicken",
      "count": 333291
    },
    {
      "recipe": "Speedy Steak Fajitas",
      "count": 333578
    },
    {
      "recipe": "Spinach Artichoke Pasta Bake",
      "count": 333545
    },
    {
      "recipe": "Steakhouse-Style New York Strip",
      "count": 333473
    },
    {
      "recipe": "Stovetop Mac 'N' Cheese",
      "count": 333098
    },
    {
      "recipe": "Sweet Apple Pork Tenderloin",
      "count": 332595
    },
    {
      "recipe": "Tex-Mex Tilapia",
      "count": 333749
    },
    {
      "recipe": "Yellow Squash Flatbreads",
      "count": 333394
    }
  ],
  "busiest_postcode": {
    "postcode": "10176",
    "delivery_count": 91785
  },
  "count_per_postcode_and_time": {
    "postcode": "10161",
    "from": "9AM",
    "to": "11PM",
    "delivery_count": 6779
  },
  "match_by_name": null
}
```
