using System;
namespace documentationExample {
	class MainClass {
  		public static void Main (string[] args) {
			Console.WriteLine ("Hello World");
  		}

  		/// <summary>
  		/// Converts inches to centimeters
  		/// </summary>
  		/// <param="increase"> The length of measurement to convert </param>
  		/// <returns> The converted value </returns>
  		private static float inchesToCentimeters(float length) {
			return length * 2.54f;
  		}
	}
}
