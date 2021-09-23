#-----------------------------------------------------------------------------
# Name:        Documentation i.e. docstrings (python.py)
# Purpose:     To provide an example of how to create docstrings in Python
#
# References: 	This program uses the NumPy/SciPy style of documentation as found
#				here: https://numpydoc.readthedocs.io/en/latest/format.html with
#				some minor modifications based on Python 3 function annotations
#				(the -> notation).
#
# Author:      Mr. Seidel
# Created:     23-Sep-2021
# Updated:     23-Sep-2021
#-----------------------------------------------------------------------------

def inchesToCentimeters(length : float) -> float:
    '''
    Converts inches to centimeters

    Parameters
    ----------
    length : float
        The length of measurement to convert

    Returns
    -------
    float
        The converted value

    '''
    return length * 2.54

print(inchesToCentimeters(1))
