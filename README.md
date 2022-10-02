Given an integer, instances, and an array, arr[] of size N representing the average utilization percentage of the computing system at each second, the task is to find the number of instances at the end of the time frame such that the computing system auto-scales the number of instances according to the following rules: 

  Average utilization < 25%: Reduce the number of instances by half if the number of instances is greater than 1.
- 25% ≤ Average utilization ≤ 60%: Take no action.
- Average utilization > 60%: Double the number of instances if the doubled value does not exceed 2* 108.
Once an action of adding or reducing the number of instances is performed, the system will stop monitoring for 10 seconds. During that time, the number of instances does not change.

Examples:

Input: instances = 2, arr[] = {25, 23, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 76, 80}
Output: 2
Explanation:
At second 1, arr[0] = 25 ≤ 25, so take no action.
At second 2, arr[1] = 23 < 25, so an action is instantiated to halve the number of instances to ceil(2/2) = 1. The system will stop checking for 10 seconds, so from arr[2] through arr[11] no actions will be taken.
At second 13, arr[12] = 76 > 60, so the number of instances is doubled from 1 to 2.
There are no more readings to consider and 2 is the final value.

Input: instances = 5, arr = {30, 5, 4, 8, 19, 89}
Output: 3
Explanation:
At second 1, 25 ≤ arr[0] = 30 ≤ 60, so take no action.
At second 2, arr[1] = 5 < 25, so an action is instantiated to halve the number of instances to ceil(5/2) = 3.
The system will stop checking for 10 seconds, so from arr[2] through arr[5] no actions will be taken.

There are no more readings to consider and 3 is the final answer.