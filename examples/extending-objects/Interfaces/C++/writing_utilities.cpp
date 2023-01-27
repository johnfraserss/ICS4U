#pragma once
#include <iostream>
#include <fstream>

/*
    An interface that all other classes in this program extends from. This interface outlines all the basic functionality any utility should have if you want to write with it.
*/
class WriteableUtility
{
protected:
    int length;

    /*
    The constructor for this interface.

    @param length: int   Specifies how long this writing utility is. This is used in the program to help identify which messages were written by what utilities in the output file.
    */
public:
    WriteableUtility(const int &num) : length(num) {}

    /*
    The write method which takes a message and outputs it. This method will be overwritten by its descendants, as the output changes according to what utility is writing the message.

    @param msg: std::string   The message you want to write.
    @param paper: std::ofstream   The file output stream that the message goes to.
    */
    virtual void write(std::string msg, std::ofstream &paper) = 0;

    /*
    This method will scribble random letters and output them. This method will be overwritten by its descendants, as the output changes according to what utility is writing the message.

    @param paper: std::ofstream   The file output stream that the scribbled message goes to.
    */
    virtual inline void scribble(std::ofstream &paper) = 0;
};


#include "pen_utilities.h"