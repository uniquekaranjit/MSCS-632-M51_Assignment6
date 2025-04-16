import java.util.LinkedList;
import java.util.Queue;
import java.util.concurrent.locks.ReentrantLock;

public class TaskQueue {
    private final Queue<String> queue = new LinkedList<>();
    private final ReentrantLock lock = new ReentrantLock();

    public void addTask(String task) {
        lock.lock();
        try {
            queue.add(task);
        } finally {
            lock.unlock();
        }
    }

    public String getTask() {
        lock.lock();
        try {
            return queue.poll(); // returns null if empty
        } finally {
            lock.unlock();
        }
    }

    public boolean isEmpty() {
        lock.lock();
        try {
            return queue.isEmpty();
        } finally {
            lock.unlock();
        }
    }
}
