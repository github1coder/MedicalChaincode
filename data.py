import pandas


# a = [1, 2, 3]
# myvar = pandas.Series(a)
# print(myvar)

# data = [['Google',10],['Runoob',12],['Wiki',13]]
# df = pandas.DataFrame(data,columns=['Site','Age'])
# print(df)

# pandas.set_option('display.max_columns', None)
# pandas.set_option('display.max_rows', 2)

csv = pandas.read_csv('log.csv')
print(csv.to_string())
# print(csv.loc[0].to_string())
# print(pandas.DataFrame())
# print("Hello World")