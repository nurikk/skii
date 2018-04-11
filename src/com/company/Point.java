package com.company;

import java.util.ArrayList;
import java.util.List;

public class Point {
    private int elevation;
    private boolean isPeak;
    private List<Point> neighbors = new ArrayList();

    public Point(int elevation) {
        this.elevation = elevation;
    }

    public int getElevation() {
        return elevation;
    }

    public boolean isPeak() {
        return isPeak;
    }

    public void setPeak(boolean peak) {
        isPeak = peak;
    }

    public List<Point> getNeighbors() {
        return neighbors;
    }




    public void addNeighbor(Point n){
        this.neighbors.add(n);
    }

    @Override
    public String toString() {
        return "elevation: " + this.getElevation();
    }
}
