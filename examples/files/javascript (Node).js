/**
 * If you are not using P5.js and are instead using Node.JS, you can use the FS (file system) library to access local files
 */

let fs = require("fs");

/**
 * read the file using the fs library - Please note that this operation is asynchronous 
 *   Asynchronous meaning: it won't wait to complete the operation before continuing through the program
 * fs.readFile returns an error and data object
 */ 
fs.readFile('input.txt', (err, data) => {
    if (err) {
        // console.error is similar to console.log, only in that it will display errors differently
        console.error(err);
    } else {
        console.log("output of file: " + data.toString());
    }
})
