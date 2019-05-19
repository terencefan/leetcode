package p207;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Set;

class Solution {

    Map<Integer, Boolean> canFinish;

    Map<Integer, List<Integer>> map = new HashMap<>();

    boolean dfs(int course, Set<Integer> visited) {
        if (visited.contains(course)) {
            return false;
        } else if (canFinish.get(course) != null) {
            return canFinish.get(course);
        } else if (!map.containsKey(course)) {
            canFinish.put(course, true);
            return true;
        }

        visited.add(course);
        for (int pre : map.get(course)) {
            if (visited.contains(pre)) {
                canFinish.put(course, false);
                return false;
            }
            if (!dfs(pre, visited)) {
                canFinish.put(course, false);
                return false;
            }
        }
        canFinish.put(course, true);
        return true;
    }

    public boolean canFinish(int numCourses, int[][] prerequisites) {
        this.canFinish = new HashMap<>();
        this.map = new HashMap<>();

        for (int[] pair : prerequisites) {
            int p = pair[0];
            int q = pair[1];
            if (!map.containsKey(p)) {
                map.put(p, new ArrayList<>());
            }
            map.get(p).add(q);
        }

        for (int i = 0; i < numCourses; i++) {
            if (!dfs(i, new HashSet<>())) {
                return false;
            }
        }
        return true;
    }
}
