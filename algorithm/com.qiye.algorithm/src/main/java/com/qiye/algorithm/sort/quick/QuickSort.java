package com.qiye.algorithm.sort.quick;

import com.qiye.algorithm.sort.Sort;

public class QuickSort extends Sort {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		int[] a = { 33, 22, 1, 4, 2, 123, 56, 888 };
		Sort sort = new QuickSort();
		sort.sort(a);
		sort.outArray(a);
	}

	@Override
	public void sort(int[] a) {
		// TODO Auto-generated method stub
		if (a == null || a.length <= 1) {
			return;
		}
		this.partition(a, 0, a.length - 1);
	}

	private void partition(int[] a, int s, int e) {
		if (s == e) {
			return;
		}
		int flag = a[s];
		int head = s;
		int end = e;
		while (head < end) {
			while (head < end && a[end] >= flag) {
				end--;
			}
			a[head] = a[end];
			while (head < end && a[head] <= flag) {
				head++;
			}
//			if(end>head){
				a[end] = a[head];
//			}
			
		}
		a[head] = flag;
		if (head - s > 1) {
			this.partition(a, s, head - 1);
		}
		if (e - head > 1) {
			this.partition(a, head + 1, e);
		}
	}

}
