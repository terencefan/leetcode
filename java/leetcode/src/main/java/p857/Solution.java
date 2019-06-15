package p857;

import java.util.Comparator;
import java.util.PriorityQueue;
import java.util.Queue;

class Solution {

	class Worker {

		int quality;
		int wage;
		double ratio;

		Worker(int quality, int wage) {
			this.quality = quality;
			this.wage = wage;
			this.ratio = (double) wage / quality;
		}
	}

	class RatioComparator implements Comparator<Worker> {

		@Override
		public int compare(Worker o1, Worker o2) {
			return Double.compare(o1.ratio, o2.ratio);
		}
	}

	class QualityComparator implements Comparator<Worker> {
		@Override
		public int compare(Worker o1, Worker o2) {
			return Integer.compare(o2.quality, o1.quality);
		}
	}

	public double mincostToHireWorkers(int[] quality, int[] wage, int K) {
		Queue<Worker> q = new PriorityQueue<>(quality.length, new RatioComparator());

		for (int i = 0; i < quality.length; i++) {
			q.offer(new Worker(quality[i], wage[i]));
		}

		Queue<Worker> selected = new PriorityQueue<>(quality.length, new QualityComparator());

		int sumQuality = 0;
		double currentRatio = 0;
		double result = Double.MAX_VALUE;

		for (int i = 0; i < quality.length; i++) {
			Worker w = q.poll();
			currentRatio = w.ratio;

			if (selected.size() >= K) {
				sumQuality -= selected.poll().quality;
			}
			selected.offer(w);
			sumQuality += w.quality;

			if (selected.size() == K) {
				result = Math.min(result, currentRatio * sumQuality);
			}
		}

		return result;
	}

}
