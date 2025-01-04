from abc import abstractmethod

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
    @abstractmethod
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
        print(df_input_data.head())  # return first 5 rows
        print(df_input_data.info)  # df structure

        print("Indexing and Selection")
        print(df_input_data["player"])  # player table DF for all the players
        print(df_input_data[df_input_data["runs"] > 10000])  # DF for player > 10K runs
        print(df_input_data.iloc[0])  # first row
        print(df_input_data.loc[df_input_data["player"] == "kohli"])  # conditional row match

        print("Add/Drop colums")
        df_input_data["player_avg"] = df_input_data["runs"] / df_input_data["matches"]
        print(df_input_data)
        # df_input_data.drop("matches", axis=1) # axis = 0, axis = 1 is colum wise drop

        print("Sorting and Ranking")
        print(df_input_data.sort_values("runs", ascending=False))  # sort by runs in desc
        print(df_input_data["runs"].rank(ascending=False))  # rank 1 to N based on the runs col

        print("Handling missing data")

        input_missing_data = {
            "player": ["sachin", "dhoni", "kohli", None],
            "runs": [10000, 9000, None, 8000]
        }

        df_missing_data = pd.DataFrame(input_missing_data)
        print(df_missing_data.isnull())  # find the missing values from DF
        df_missing_data["runs"] = df_missing_data["runs"].fillna(0)  # fill the runs as NaN to zero
        print(df_missing_data)
        df_missing_data = df_missing_data.dropna()  # drop the NaN
        print(df_missing_data)

        print("Merging, Joining and concatenation")
        df1 = pd.DataFrame({"player": ["sachin", "dhoni"], "runs": [10000, 9000]})
        df2 = pd.DataFrame({"player": ["sachin", "dhoni"], "average": [44.83, 50.57]})
        merge_df = pd.merge(df1, df2, on="player")
        print(merge_df)
        pd_concat = pd.concat([df1, df2])  # will concat with missing values as NaN
        print(pd_concat)

        print("Pivot tables")
        pivot_data = {
            "player": ["sachin", "dhoni", "sachin", "dhoni"],
            "year": [2003, 2003, 2011, 2011],
            "runs": [673, 400, 482, 350]
        }

        # pivot table for col year and player as index and output the runs with aggredate function to sum the runs
        df_pivot = pd.DataFrame(pivot_data)
        pd_pivot = df_pivot.pivot_table(columns="year", index="player", values="runs", aggfunc="sum")
        print(pd_pivot)

        print("R+W from csv")
        # df_csv = pd.read_csv("example.csv")
        # df_csv.to_csv("example.csv", index=False) # donot index the first column
        # print(df_csv)

        print("Custom functions")
        df_input_data["strike_rate"] = df_input_data["runs"].apply(
            lambda x: x / 50 if x else 0
        )

        print("Time series")
        date_range = pd.date_range("2025-01-04", periods=5)
        df_date_range = pd.DataFrame({"date": date_range, "runs": [10, 20, 100, 51, 100]})
        print(df_date_range)


class Libraries(Sample):
    @abstractmethod
    def execute(self):
        # NumPy().execute()
        Pandas().execute()
