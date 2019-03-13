def minimum_index(seq):
    if len(seq) == 0:
        raise ValueError("Cannot get the minimum value index from an empty sequence")
    min_idx = 0
    for i in range(1, len(seq)):
        if seq[i] < seq[min_idx]:
            min_idx = i
    return min_idx
import random

class TestDataEmptyArray(object):
    
    @staticmethod
    def get_array():
        # complete this function
        return []

class TestDataUniqueValues(object):

    minindex_obj = 0
    minval_obj = 0
    
    @staticmethod
    def get_array():
        # complete this function
        array = []
        ct = 0
        size = random.randint(2,1000)
        for i in range(size):
            num = random.randint(-10000,10000)
            if num not in array:
                if i==0:
                    TestDataUniqueValues.minval_obj = num
                else:
                    if TestDataUniqueValues.minval_obj>num:
                        TestDataUniqueValues.minval_obj = num
                        TestDataUniqueValues.minindex_obj = ct
                array.append(num)
                ct += 1
        return array
    
    @staticmethod
    def get_expected_result():
        return TestDataUniqueValues.minindex_obj
        # complete this function

class TestDataExactlyTwoDifferentMinimums(object):
    
    @staticmethod
    def get_array():
        # complete this function
        TestDataExactlyTwoDifferentMinimums.array = []
        ct = 0
        TestDataExactlyTwoDifferentMinimums.minindex = 0
        size = random.randint(2,1000)
        for i in range(size):
            num = random.randint(-10000,10000)
            if num not in TestDataExactlyTwoDifferentMinimums.array:
                if i==0:
                    TestDataExactlyTwoDifferentMinimums.minval = num
                else:
                    if TestDataExactlyTwoDifferentMinimums.minval>num:
                        TestDataExactlyTwoDifferentMinimums.minval = num
                        TestDataExactlyTwoDifferentMinimums.minindex = ct
                TestDataExactlyTwoDifferentMinimums.array.append(num)
                ct += 1
        ins_point = random.randint(0,ct)
        TestDataExactlyTwoDifferentMinimums.array.insert(ins_point,TestDataExactlyTwoDifferentMinimums.minval)
        if ins_point < TestDataExactlyTwoDifferentMinimums.minindex:
            TestDataExactlyTwoDifferentMinimums.minindex = ins_point
        return TestDataExactlyTwoDifferentMinimums.array
    
    @staticmethod
    def get_expected_result():
        return TestDataExactlyTwoDifferentMinimums.minindex
        # complete this function


def TestWithEmptyArray():
    try:
        seq = TestDataEmptyArray.get_array()
        result = minimum_index(seq)
    except ValueError as e:
        pass
    else:
        assert False


def TestWithUniqueValues():
    seq = TestDataUniqueValues.get_array()
    assert len(seq) >= 2

    assert len(list(set(seq))) == len(seq)

    expected_result = TestDataUniqueValues.get_expected_result()
    result = minimum_index(seq)
    assert result == expected_result


def TestiWithExactyTwoDifferentMinimums():
    seq = TestDataExactlyTwoDifferentMinimums.get_array()
    assert len(seq) >= 2
    tmp = sorted(seq)
    assert tmp[0] == tmp[1] and (len(tmp) == 2 or tmp[1] < tmp[2])

    expected_result = TestDataExactlyTwoDifferentMinimums.get_expected_result()
    result = minimum_index(seq)
    assert result == expected_result

TestWithEmptyArray()
TestWithUniqueValues()
TestiWithExactyTwoDifferentMinimums()
print("OK")

