package com.qiye.algorithm.sort.merge;

import com.qiye.algorithm.sort.Sort;

public class ArrayMergeSort extends Sort{

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		int[] a = { 33, 22, 1, 4, 2, 123, 56, 888 };
		Sort sort = new ArrayMergeSort();
		sort.sort(a);
		sort.outArray(a);
	}

	public void sort(int[] a) {
		if (a == null || a.length <= 1) {
			return;
		}
		this.mergeSort(a, 0, a.length-1);
	}

	private void mergeSort(int[] a, int s, int e) {
		if (s >= e) {
			return;
		} else {
			mergeSort(a, s, (s + e) / 2);
			mergeSort(a, (s + e) / 2 + 1, e);
			merge(a, s, (s + e) / 2, (s + e) / 2 + 1, e);
		}
	}

	private void merge(int[] a, int s1, int e1, int s2, int e2) {
		int[] tmp = new int[e2 - s1 + 1];
		int i = s1, j = s2;
		int index = 0;
		while (i <= e1 && j <= e2) {
			if (a[i] <= a[j]) {
				tmp[index] = a[i];
				i++;
			} else {
				tmp[index] = a[j];
				j++;
			}
			index++;
		}
		while (i <= e1) {
			tmp[index++] = a[i++];
		}
		while (j <= e2) {
			tmp[index++] = a[j++];
		}
		for (i = 0; i < tmp.length; i++) {
			a[s1++] = tmp[i];
		}
	}

	

}
