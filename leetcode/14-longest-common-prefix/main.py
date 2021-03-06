"""
    2nd approach:
	- instead of appending charactors so many times, slice the end
    
	Time		O(n*m)
	Space		O(m)
	0ms beats 100%
	2jun2019
"""


class Solution(object):
    def longestCommonPrefix(self, strs):
        """
        :type strs: List[str]
        :rtype: str
        """
        if len(strs) == 0:
            return ""
        res = strs[0]
        for i in range(1, len(strs)):
            s = strs[i]
            j = 0
            while j < len(s):
                if j >= len(res) or s[j] != res[j]:
                    break
                j += 1
            res = res[:j]
        return res
