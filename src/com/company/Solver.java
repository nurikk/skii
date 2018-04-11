
package com.company;


import java.io.FileReader;
import java.io.IOException;
import java.io.LineNumberReader;
import java.util.*;
import java.util.function.Function;
import java.util.function.Predicate;
import java.util.stream.Collectors;

/**
 * Created by ainur.timerbaev on 10/4/18.
 */

public class Solver {
    private String fileName;


    private int[][] field;
    private int maxRows;
    private int maxCols;
    private List<List<Integer>> results = new ArrayList<>();


    /**
     * @param fileName ski map file
     */
    Solver(String fileName) {
        this.fileName = fileName;
    }
    private void readFile(String fileName) {
        try {
            LineNumberReader lineNumberReader = new LineNumberReader(new FileReader(fileName));
            String line;
            while ((line = lineNumberReader.readLine()) != null) {
                int lineNumber = lineNumberReader.getLineNumber() - 2;
                int[] lineData = Arrays.stream(line.split(" ")).mapToInt(Integer::parseInt).toArray();
                if (lineNumber == -1) {
                    this.maxRows = lineData[0];
                    this.maxCols = lineData[1];
                    this.field = new int[this.maxRows][this.maxCols];
                } else {
                    this.field[lineNumber] = lineData;
                }
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
    private boolean exists(int row, int cell) {
        return row >= 0 && row < this.maxRows && cell >= 0 && cell < this.maxCols;
    }

    private List<int[]> getNeighbors(int row, int cell) {
        List<int[]> neighbors = new ArrayList<>();

       int[][] directions = new int[][]{
               {-1, 0},
               {0, -1},
               {1, 0},
               {0, 1}
       };

        for (int[] direction : directions) {
            int moveX = row + direction[0];
            int moveY = cell + direction[1];
            if (this.exists(moveX, moveY)) {
                neighbors.add(new int[]{moveX, moveY});
            }
        }
       return neighbors;
    }

    private List<int[]> findPeaks() {
        List<int[]> peaks = new ArrayList<>();

        for (int row = 0; row < this.maxRows; row++){
            for(int coll = 0; coll < this.maxCols; coll++){
                int finalRow = row;
                int finalColl = coll;
                Predicate<int[]> check = neighbor -> this.field[finalRow][finalColl] > this.field[neighbor[0]][neighbor[1]];

                List<int[]> neighbors = this.getNeighbors(row, coll);
                if (neighbors.stream().allMatch(check)){
                    peaks.add(new int[]{finalRow, finalColl});
                }
            }
        }
        return peaks;
    }

    private List<int[]> getPossibleDirections(int row, int col){
        Predicate<int[]> check = neighbor -> this.field[row][col] > this.field[neighbor[0]][neighbor[1]];

        return this.getNeighbors(row, col).stream().filter(check).collect(Collectors.toCollection(ArrayList::new));
    }


    private void findNextSteps(int row, int col, List<Integer> currentPath) {
        currentPath.add(this.field[row][col]);

        List<int[]> movements = this.getPossibleDirections(row, col);

        movements.forEach(movement -> {
            List<Integer> copyCurrentPath = new ArrayList<>(currentPath);
            List<int[]> nextMovements = this.getPossibleDirections(movement[0], movement[1]);
            if (nextMovements.isEmpty()) {
                copyCurrentPath.add(this.field[movement[0]][movement[1]]);
                this.results.add(copyCurrentPath);
            } else {
                this.findNextSteps(movement[0], movement[1], copyCurrentPath);
            }

        });
    }
    public void solve(){

        this.readFile(this.fileName);
        List<int[]> peaks = this.findPeaks();
        peaks.forEach(peak -> this.findNextSteps(peak[0], peak[1], new ArrayList<>()));

        Function<List<Integer>, Integer> getDrop = (path) -> path.get(0) - path.get(path.size() - 1);
        Comparator<List<Integer>> comparator = Comparator.comparingInt(List::size);
        comparator = comparator.thenComparing(Comparator.comparing(getDrop));

        this.results.sort(comparator);

        List<Integer> winner =  this.results.get(this.results.size() - 1);
        System.out.println("winner:" + winner);
        System.out.println("len:" + winner.size());
        System.out.println("drop:" + getDrop.apply(winner));

    }
}
