//Example from https://www.dotnetperls.com/file-open

using System;
using System.IO;

class Program
{
    static void Main()
    {
        using (FileStream stream = File.Open("C:\\bin", FileMode.Create))
        using (BinaryWriter writer = new BinaryWriter(stream))
        {
            writer.Write(303);
            writer.Write(720);
        }
        using (FileStream stream = File.Open("C:\\bin", FileMode.Open))
        using (BinaryReader reader = new BinaryReader(stream))
        {
            int a = reader.ReadInt32();
            int b = reader.ReadInt32();

            Console.WriteLine(a);
            Console.WriteLine(b);
        }
    }
}