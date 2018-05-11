function wait(f, cond, ...params) {
    if (cond()) {
        f(...params);
    } else {
        setTimeout(function() {
            wait(f, cond, ...params);
        }, 5)
    }
}

class Philosopher {
    constructor(id, pc) {
        this.id = id;
        this.pc = pc;
    }
    signalWake() {
        this.signal('philosopherWakes');
    }
    signalEat() {
        this.signal('philosopherEats');
    }
    signalSleep() {
        this.signal('philosopherSleeps');
    }
    // eat() {
    //     setTimeout(() => this.signalSleep(), 250);
    // }
    wake() {
        console.log('Waking in 10:', this.id);
        setTimeout(() => this.signalWake(), 10);
    }
    sleep() {
        let rand = Math.floor( Math.random() * 1250 );
        console.log('Sleeping in', rand, ':', this.id);
        setTimeout(() => this.signalSleep(), rand);
    }
    signal(action) {
        this.pc.respondToAction(this.id, action);
    }
}


class PhilosopherController {
    constructor(numP) {
        // numP = 4;
        this.philosophers = Array.apply(null, { length: numP }).map( (_, i) => new Philosopher(i, this) );
        this.eating = [Array.apply(null, { length: numP }).map(Boolean)];
        this.forks = Array.apply(null, { length: numP + 1 }).map(() => true);
    }
    getBoundedIndex(num) { return num % this.philosophers.length}
    getStatus(id) {
        let philosopher = this.philosophers[id];
        let isEating = this.eating[id];
        let forks = {l: this.forks[this.getBoundedIndex(id + 1)], r: this.forks[id]};
        return {philosopher, isEating, forks};
    }
    respondToAction(id, action) {
        console.log(id);
        this[action].bind(this)(id);

    }
    philosopherWakes(id) {
        let { philosopher,
            isEating,
            forks
        } = this.getStatus(id);
        let doTest = () => {
            let { forks } = this.getStatus(id);
            return forks.l && forks.r;
        }
        if (doTest()) {
            console.log('eating:', id);
            this.eating[id] = true;
            this.forks[id] = false;
            this.forks[this.getBoundedIndex(id + 1) ] = false;
            philosopher.sleep();
        } else {
            wait(this.philosopherWakes, doTest, id);
        }
    }
    philosopherSleeps(id) {
        console.log(id);
        this.forks[id] = true;
        this.forks[this.getBoundedIndex(id + 1)] = true;
        this.eating[id] = false;
        this.philosophers[id].wake();
    }
    start() {
        this.philosophers
            .forEach((p) => 
                setTimeout( () => p.wake(),
                Math.random() * 500))
    }
}

let pc = new PhilosopherController(3)
pc.start();