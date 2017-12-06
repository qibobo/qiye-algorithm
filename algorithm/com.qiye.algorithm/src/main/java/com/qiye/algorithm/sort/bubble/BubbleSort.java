package com.qiye.algorithm.sort.bubble;

import com.qiye.algorithm.sort.Sort;

public class BubbleSort extends Sort {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		int[] a = { 33, 22, 1, 4, 2, 123, 56, 888 };
		Sort sort = new BubbleSort();
		sort.sort(a);
		sort.outArray(a);
	}

	@Override
	public void sort(int[] a) {
		// TODO Auto-generated method stub
		for (int i = a.length-1; i >= 0; i--) {
			for (int j = i; j >= 0; j--) {
				if (a[i] < a[j]) {
					int tmp = a[i];
					a[i] = a[j];
					a[j] = tmp;
				}
			}
		}
	}

}
