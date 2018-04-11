
package com.company;


import java.io.FileReader;
import java.io.IOException;
import java.io.LineNumberReader;
import java.util.*;
import java.util.function.Function;
import java.util.stream.Collectors;

/**
 * Created by ainur.timerbaev on 10/4/18.
 */

public class Solver {
    private String fileName;


    private List<List<Point>> field = new ArrayList<>();
    private int maxRows;
    private int maxCols;
    private List<List<Integer>> results = new ArrayList<>();

    private int maxPathLen = 0;

    /**
     * @param fileName ski map file
     */
    Solver(String fileName) {
        this.fileName = fileName;
    }
    private void readFile(String fileName) {
        long start = System.currentTimeMillis();
        try {
            LineNumberReader lineNumberReader = new LineNumberReader(new FileReader(fileName));
            String line;
            while ((line = lineNumberReader.readLine()) != null) {
                int lineNumber = lineNumberReader.getLineNumber() - 2;
                int[] lineData = Arrays.stream(line.split(" ")).mapToInt(Integer::parseInt).toArray();
                if (lineNumber == -1) {
                    this.maxRows = lineData[0];
                    this.maxCols = lineData[1];
                } else {
                    List<Point> newRow = Arrays.stream(lineData).mapToObj(elevation -> new Point(elevation)).collect(Collectors.toList());
                    this.field.add(lineNumber, newRow);

                }
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
        System.out.printf("readFile:  %d ms%n", (System.currentTimeMillis() - start));
    }

    private boolean exists(int row, int cell) {
        return row >= 0 && row < this.maxRows && cell >= 0 && cell < this.maxCols;
    }

    private void findNextSteps(Point peak, List<Integer> currentPath) {
        currentPath.add(peak.getElevation());

        peak.getNeighbors().forEach(movement -> {
            List<Integer> copyCurrentPath = new ArrayList<>(currentPath);
            if (movement.getNeighbors().isEmpty()) {
                copyCurrentPath.add(movement.getElevation());
                if (copyCurrentPath.size() >= this.maxPathLen){
                    this.results.add(copyCurrentPath);
                    this.maxPathLen = copyCurrentPath.size();
                }
            } else {
                this.findNextSteps(movement, copyCurrentPath);
            }

        });
    }
    private void buildMap(){
        long start = System.currentTimeMillis();
        int[][] directions = new int[][]{
                {-1, 0},
                {0, -1},
                {1, 0},
                {0, 1}
        };
       for (int row = 0; row < this.maxRows; row++){
           for (int col = 0; col < this.maxCols; col++){
               boolean isPeak = true;
               Point currPoint = this.field.get(row).get(col);
               for (int[] direction : directions) {
                   int moveX = row + direction[0];
                   int moveY = col + direction[1];
                   if (this.exists(moveX, moveY)) {
                       Point n = field.get(moveX).get(moveY);
                       if (currPoint.getElevation() <= n.getElevation()){
                           isPeak = false;
                       } else {
                           currPoint.addNeighbor(n);
                       }
                   }
               }
               currPoint.setPeak(isPeak);
           }
       }
        System.out.printf("buildMap:  %d ms%n", (System.currentTimeMillis() - start));
    }
    public void solve(){
        long start = System.currentTimeMillis();
        this.readFile(this.fileName);
        this.buildMap();

        long startIterate = System.currentTimeMillis();
        for (List<Point> row : this.field) {
            row.stream().filter(p -> p.isPeak()).forEach(peak -> {
                this.findNextSteps(peak, new ArrayList<>());
            });
        }
        System.out.printf("iterate findNextSteps:  %d ms%n", (System.currentTimeMillis() - startIterate));


        long startFinishing = System.currentTimeMillis();
        Function<List<Integer>, Integer> getDrop = (path) -> path.get(0) - path.get(path.size() - 1);
        Comparator<List<Integer>> comparator = Comparator.comparingInt(List::size);
        comparator = comparator.thenComparing(Comparator.comparing(getDrop));
        this.results.sort(comparator);
        List<Integer> winner =  this.results.get(this.results.size() - 1);
        System.out.printf("startFinishing:  %d ms%n", (System.currentTimeMillis() - startFinishing));

        System.out.println("winner:" + winner);
        System.out.println("len:" + winner.size());
        System.out.println("drop:" + getDrop.apply(winner));
        System.out.printf("solve:  %d ms%n", (System.currentTimeMillis() - start));

    }
}
