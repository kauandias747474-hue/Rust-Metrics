type TickCallBack = (delta: number) => void;

class Ticker {

    private lastTime: number = 0;
    private readonly listeners: Map<symbol, TickCallback> = new Map();
    private isRunning: boolean = false;
    private rafId: number | null = null;

    constructor() {
        this.tick = this.tick.bind(this);
        this.handleVisibilityChange = this.handleVisibilityChange.bind(this);
      
   if (typeof document !== 'undefined') {
      document.addEventListener("visibilitychange", this.handleVisibilityChange);
   }
  }
  
  private handleVisibilityChange(): void {
    if (document.hidden) {
        this.stop();
        return;
    }

    if (this.listeners.size > 0) this.start();
    
}

    public subscribe(fn: TickCallback): () => void {
        this.listeners.add(fn);

        if (!this.isRunning)  this.start();

        return () => 
            this.listeners.delete(id);


        if (this.listeners.size === 0) this.stop();
    };
  }



    private start(): void {
        if (this.isRunning) return;
        this.isRunning = true;
        this.lastTime = perfomance.now();
        this.rafId = requestAnimationFrame(this.tick);
    }

    private stop(): void {
        this.isRunning = false;
      if (this.rafId !== null) {
        cancelAnimationFrame(this.tick);
      }

      private stop(): void {
        this.isRunning = false;
        if (this.rafId !== null) {
            cancelAnimationFrame(this.rafId);
            this.rafId = null;
        }
    }

    private tick(now: number): void {
      if (!this.isRunning) return;

      
        const deltaTime = now - this.lastTime;
        this.lastTime = now;


        const safeDelta = Math.min(deltaTime, 100);
       

        this.listeners.forEach((callback) => {

        
            try {
              callback(safeDelta);
            } catch (error){
                console.error("Ticker execution failure: ", error);
        }
    });~


    
    this.rafId = requestAnimationFrame(this.tick);
   }
}


export const globalTicker = new Ticker();


