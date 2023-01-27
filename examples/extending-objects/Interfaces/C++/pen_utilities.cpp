#pragma once
#include <vector>
#include "writing_utilities.h"

/*
    A Pen object that you can write with, which implements WriteableUtility with several customizations and methods.
*/
class Pen : public WriteableUtility
{
private:
    bool capped = true;
    std::string engraving;

public:

    /*
        The constructor for a Pen object. Depending on what parameters you give this Pen object, that will greatly change how outputted messages will look like.

        @param len: int   The length of the pen. This is used to help identify which messages are written with it in the output.
        @param engraving: std::string   The engraving on the pen. This engraving will show up before and after every written message in the output.
    */
    Pen(const int &len, const std::string &engraving) : WriteableUtility(len), engraving(engraving) {}

    /*
        The write method for a Pen object that overrides WriteableUtility's. This takes the message provided to it and outputs it on the stream provided, with many modifications based on the object's attributes which are outlined in detail above. Please note that a message will only be written if the Pen is currently uncapped (capped = false).

        @param msg: std::string   The message that you want to write.
        @param paper: std::ofstream   The stream that you want to output to.
    */
    void write(std::string msg, std::ofstream &paper) override
    {
	// Fully capitalize input message
        std::transform(msg.begin(), msg.end(), msg.begin() ::toupper);
        if (!capped)
        {
            paper << "\n"
                  << this->engraving << " " << msg
                  << " " << this->engraving << "\n(Pen length: " << this->length << ")" << std::endl;
        }
    }

    /*
        The scribble method for a Pen object that overrides WriteableUtility's. This outputs a random message on the stream provided, with many modifications based on the object's attributes which are outlined in detail above. Please note that a message will only be written if the Pen is currently uncapped (capped = false).

        @param paper: std::ofstream   The stream that you want to output to.
    */
    inline void scribble(std::ofstream &paper) override
    {
        if (!capped)
        {
            paper << "\n"
                  << this->engraving << " "
                  << "Random words are being scribbled..."
                  << " " << this->engraving << "\n(Pen length: " << this->length << ")" << std::endl;
        }
    }

    /*
        As in real life, this simple method "clicks" the pen flipping its state between being capped and uncapped. Pens first must be uncapped (capped = false) before they can be written with.
    */
    inline void click() { this->capped = !this->capped; }
};
