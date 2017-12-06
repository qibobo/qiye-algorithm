package com.qiye.algorithm.sort.insert;

import com.qiye.algorithm.sort.Sort;

public class ArrayInsertSort  extends Sort{

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		int[] a = {3,22,1,4,2,123,56,888};
		Sort sort = new ArrayInsertSort();
		sort.sort(a);
		sort.outArray(a);
	}

	// asc
	public void sort(int[] a) {
		if (a == null || a.length <= 1) {
			return;
		}
		int i, j;
		for (i = 1; i < a.length; i++) {
			int tmp = a[i];
			for (j = 0; j < i; j++) {
				if (a[j] > a[i]) {
					break;
				}
			}
			for (int m = i - 1; m >= j; m--) {
				a[m + 1] = a[m];
			}
			a[j] = tmp;
		}
	}
	

}
