import numpy as np
import pandas as pd


class Sample:
    def execute(self):
        pass


# Numerical Computation
class NumPy(Sample):
    def execute(self):
        num_py_array = np.array([3, 1, 2, 9, 8])

        print("Slice")
        print(num_py_array)
        print(num_py_array[1:3])  # suffix exclusive and prefix inclusive index print [1,2]
        print(num_py_array[num_py_array >= 8])  # conditional

        print("Aggregate")
        print(num_py_array.sum())
        print(num_py_array.mean())
        print(num_py_array.min())
        print(num_py_array.max())
        num_py_array.sort()
        print(num_py_array)

        print("Reshaping")
        shape = np.arange(1, 4)  # array of [1 to 3]
        print(shape)
        reshape = shape.reshape((3, 1))
        print(reshape)
        print(reshape.flatten())

        print("2D Array")
        num_py_array_2d = np.array([[1, 2, 4], [4, 5, 6]])
        print(num_py_array_2d)

        ones_2cross3 = np.ones((2, 3))  # in shape of 2 * 3
        print(ones_2cross3)

        zeros_3cross4 = np.zeros((3, 4))  # in shape of 3 * 4
        print(zeros_3cross4)

        array_dot1 = np.array([1, 3, 4])
        array_dot2 = np.array([2, 3, 4])
        print(np.dot(array_dot1, array_dot2))

        print("Stacking V and H")
        stack_array1 = np.array([1, 4, 5])
        stack_array2 = np.array([6, 7, 8])
        h_stack = np.hstack((stack_array1, stack_array2))
        print(h_stack)  # 1, 4, 5, 6, 7, 8
        v_stack = np.vstack((stack_array1, stack_array2))
        print(v_stack)

        print("Random")
        rand = np.random.rand(2, 4)  # in shape of 2 * 4 unique random
        print(rand)

        rand_g = np.random.randn(2, 4)  # Gaussian distribution random func
        print(rand_g)


# Data manipulations
class Pandas(Sample):
    def execute(self):
        # create a data frame
        input_data = {
            "player": ["sachin", "dhoni", "kohli", "rohit"],
            "runs": [10000, 9000, 12000, 9500],
            "matches": [463, 350, 274, 243]
        }

        df_input_data = pd.DataFrame(input_data)
        print(df_input_data)

        print("DF infos")
        print(df_input_data.describe(percentiles=[1.0]))
        print(df_input_data.head()) # return first 5 rows
        print(df_input_data.info) # df structure


class Libraries(Sample):
    def execute(self):
        # NumPy().execute()
        Pandas().execute()
