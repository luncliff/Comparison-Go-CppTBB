
long SerialFib(long n) {
    if (n < 2)
        return n;
    else
        return SerialFib(n - 1) + SerialFib(n - 2);
}



class FiboTask : public task {

public:
    const long n;
    long* const sum;

    FiboTask(long n_, long* sum_) :
        n(n_), sum(sum_)
    {
        //std::cout << this->ref_count() << "\n";

        //std::this_thread::get_id();
        //std::cout << std::this_thread::get_id() << ": ctor" << std::endl;
    }

    // Overrides virtual function task::execute
    task* execute() override 
    {      
        //std::cout << std::this_thread::get_id() << "\n";

        if (n < 2) {
            *sum = SerialFib(n);
        }
        else {
            long x, y;
            FiboTask* a = new(allocate_child()) FiboTask(n - 1, &x);
            FiboTask* b = new(allocate_child()) FiboTask(n - 2, &y);
            // Set ref_count to 'two children plus one for the wait".
            set_ref_count(2);
            // Start b running.
            spawn(*b);
            wait_for_all();
            // Start a running and wait for all children (a and b).
            //spawn_and_wait_for_all(a);
            // Do the sum
            *sum = x + y;
            return a;
        }
        //std::cout << this->ref_count() << "\n";
        return nullptr;
    }
};


long ParallelFib(long n) {
    long sum;
    FiboTask& a = *new(task::allocate_root()) FiboTask(n, &sum);
    task::spawn_root_and_wait(a);
    return sum;
}


int main() {
    //std::cout << std::thread::hardware_concurrency() << std::endl;;
    ParallelFib(3);
    //std::cout << SerialFib(4) << "\n";

    return std::system("pause");
}