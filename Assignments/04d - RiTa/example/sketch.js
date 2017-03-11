var txt;
var textArray;
var rita_txt;
var rita_pos;
var nn;

function preload() {
  doneLoading = false;
  textArray = loadStrings("euripides-bacchae-short.txt", doneLoadingText);
}

function doneLoadingText() {
  txt = textArray.join("\n");
  doneLoading = true;
  rita_txt = RiString(txt);
  rita_pos = rita_txt.pos();
}

function setup() {
  noCanvas();
  nn = createCheckbox('nn', false)
  nn.changed(changeNN);
}

function changeNN() {
  if (this.checked()) {
    for (var i = 0; i < rita_txt.length(); i++) {
      var s = createSpan(rita_txt.words()[i] + " ");
      if (rita_pos[i].charAt(0) === 'n') { //means i'm a noun of some sort
        s.style("background-color: #99DD99");
      }
    }
  }
}
