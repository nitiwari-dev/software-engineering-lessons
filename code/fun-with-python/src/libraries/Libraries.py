from abc import abstractmethod

import numpy as np
import pandas as pd
import matplotlib.pyplot as plt


class Sample:
    def execute(self):
        pass


# Numerical Computation
class NumPy(Sample):

    def __init__(self):
        print("Numpy is majorly used for numerical computation")

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

    def __init__(self):
        print("Pandas are majorly used for data manipulations")

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


# Data Visualization
class MatplotLib:

    def __init__(self):
        print("Matplotlib is majorly used for Data Visualization")

    @abstractmethod
    def execute(self):
        print("Basic line plot")
        x_axis = [x for x in range(1, 5)]
        y_axis = [y for y in range(2, 10, 2)]
        plt.plot(x_axis, y_axis, color="blue", marker="x", linestyle="--")
        plt.title("basic line plot")
        plt.xlabel("x axis")
        plt.ylabel("y axis")
        plt.grid()
        plt.show()

        print("Bar plot")
        x_bar = [x for x in range(1, 5, 1)]
        y_bar = [x for x in range(10, 50, 10)]
        plt.bar(x_bar, y_bar)
        plt.grid()
        plt.show()

        print("Horizontal Bar Plot")
        x_bar_h = [10, 30, 20]  # 10, 20, 30, 40, 50 with step of 10
        y_bar_h = [1, 2, 3]
        plt.barh(x_bar_h, y_bar_h)
        plt.show()

        print("Scattered plot")
        x_sc = [1, 2, 3, 4, 5]
        y_sc = [2, 3, 4, 7, 11]

        plt.scatter(x_sc, y_sc, color="purple", alpha=0.8, s=100)  # s is size of dot
        plt.grid()
        plt.show()

        print("Histogram")
        x_hist = np.random.randn(1000)
        bins = 20  # no. of intervals
        plt.hist(x=x_hist, bins=bins)  # out of 1000 random distribution with 20 internals
        plt.show()

        print("Pie Char")
        x_pie = [40, 20, 50, 30]
        x_labels = ["IN", "US", "UK", "JP"]
        x_explode = [0.1, 0, 0, 0]  # highlight first slice
        x_shadow = True
        x_start_angle = 90
        plt.pie(x_pie, labels=x_labels, explode=x_explode, shadow=x_shadow, startangle=x_start_angle)
        plt.show()

        print("Multiple Subplots")
        x = [1, 2, 3, 4, 5]
        y1 = [1, 4, 9, 16, 25]
        y2 = [1, 8, 27, 64, 125]

        plt.figure(figsize=(10, 5))  # 2:1 figure

        plt.subplot(1, 2, 1)  # 1 row, 2 columns, 1st subplot
        plt.plot(x, y1, color='blue')
        plt.title("Square Numbers")

        plt.subplot(1, 2, 2)  # 1 row, 2 columns, 2nd subplot
        plt.plot(x, y2, color='red')
        plt.title("Cubic Numbers")
        plt.show()

        print("Custom plot style")
        plt.style.use("ggplot")

        x = [1, 2, 3, 4, 5]
        y = [10, 20, 15, 25, 30]

        plt.plot(x, y, color='blue', marker='o')
        plt.title("Custom plot style")
        plt.xlabel("x axis")
        plt.ylabel("y axis")
        plt.grid()
        plt.show()


class Libraries(Sample):
    @abstractmethod
    def execute(self):
        # NumPy().execute()
        # Pandas().execute()
        MatplotLib().execute()
