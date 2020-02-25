'''
  Input Example
'''

f_input = open('filename.txt', 'r')
file = f_input.readlines()    # this reads in all the files to create a list called "file"
f_input.close()   # this is to close off the file when you are done reading in the data

for line in file:
  print(line)




'''
  Output Example
'''

output_string = 'This is line 1\n'    # the \n character is like hitting the enter key
output_string += 'This is line 2\n'

f_output = open('filename.txt', 'w') # w is used to overwrite a file.  You can use 'a' to append to a file instead
f_output.write(output_string) # outputs the string to the text file.  
f_output.close() # close the file for use by others
