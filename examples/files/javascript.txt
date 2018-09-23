/*
When opening a file in JavaScript, you will need to run your program using a server, and not just running in the browser.

To do so, it is suggested that you use Chrome and add the Chrome Server Extension found here:

https://chrome.google.com/webstore/detail/web-server-for-chrome/ofhbbkphhbklhfoeikjpcbhemlocgigb?hl=en

*/

//Example from https://p5js.org/reference/#/p5/loadStrings


var result;
function preload() {
  result = loadStrings('assets/test.txt');
}

function setup() {
  background(200);
  var ind = floor(random(result.length));
  text(result[ind], 10, 10, 80, 80);
}