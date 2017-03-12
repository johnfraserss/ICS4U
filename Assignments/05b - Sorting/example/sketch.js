var txt;
var textArray;

var findNames = /\b[A-Z]{5,}\b/g; //grabbing all capital lettered words that are 5 letters or more
var names;

function preload() {
  doneLoading = false;
  textArray = loadStrings("euripides-bacchae.txt", doneLoadingText);
}

function doneLoadingText() {
  txt = textArray.join("\n");
  doneLoading = true;
}

function setup() {
  noCanvas();
  names = txt.match(findNames);
  names = removeDuplicates(names);
  createP("Pre-sort");
  createP(names);

  names = insertionSort(names);

  createP("Post-sort");
  createP(names);
  for (var i = 0; i < names.length; i++) {
    //createP(names[i]);
  }
}

function draw() {

}

function removeDuplicates(n) {
  return n.filter(function(elem, index, self) {
    return index == self.indexOf(elem);
  })
}

function insertionSort(n) {
  //this is where you put in your sorting algorithm
  //remove the line below, and create your own
  //insertion sort algorithm
  //this is just for example purposes
  return n.sort();
}
