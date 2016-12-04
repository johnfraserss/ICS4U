JSONArray jsonArray;

void setup() {
  String apiStart = "http://api.wordnik.com:80/v4/word.json/";
  String word = "color";
  String apiEnd = "/definitions?limit=200&includeRelated=true&useCanonical=false&includeTags=false&";
  String apiKey = "api_key=a2a73e7b926c924fad7001ca3111acd55af2ffabf50eb4ae5";

  jsonArray = loadJSONArray(apiStart + word + apiEnd + apiKey);

  for (int i = 0; i < jsonArray.size(); i++) {
    JSONObject definition = jsonArray.getJSONObject(i);

    int sequence = definition.getInt("sequence");
    String text = definition.getString("text");
    String partOfSpeech = definition.getString("partOfSpeech");

    println(sequence + ", " + text + ", " + partOfSpeech);
  }
}