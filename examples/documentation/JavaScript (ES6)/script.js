/**
  * Converts inches to centimeters
  * 
  * @param {number} length - The length of measurement to convert
  *
  * @returns {number} The converted value
  */
function inchesToCentimeters(length) {
    return length * 2.54;
}

console.log(inchesToCentimeters(1));