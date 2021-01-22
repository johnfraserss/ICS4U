'''
  Input Example
'''

#CONVENTIONAL WAY OF READING FILES
with open('filename.txt','r') as textFile:
    file_content = textFile.readlines() # this reads in all the files to create a list called "file_content"
	# Closes file automatically


#ALTERNATE WAY OF READING FILES
'''
 f_input = open('filename.txt', 'r')
 file = f_input.readlines() # this reads in all the files to create a list called "file"   
 f_input.close()   # this is to close off the file when you are done reading in the data

'''

for line in file_content:
    print(line)






'''
  Output Example
'''

output_string = 'This is line 1\n'    # the \n character is like hitting the enter key
output_string += 'This is line 2\n'

#CONVENTIONAL WAY OF WRITING TO FILES 
'''
with open('filename.txt','w') as outfile : # w is used to overwrite a file.  You can use 'a' to append to a file instead
    outfile.write(output_string) # outputs the string to the text file. 
'''


#ALTERNATE WAY OF WRITING TO FILES 
'''
f_output = open('filename.txt', 'w') # w is used to overwrite a file.  You can use 'a' to append to a file instead
f_output.write(output_string) # outputs the string to the text file.  
f_output.close() # close the file for use by others

'''
