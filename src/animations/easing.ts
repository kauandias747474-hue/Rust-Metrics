export class AnimationMath { 

    private static signalBuffer: number[] = [];
    private static readonly MAX_BUFFER_SIZE = 10;


    
    static smoothSignal(newValue: number): number {
        this.signalBuffer.push(newValue);
        if (this.signalBuffer.length > this.MAX_BUFFER_SIZE) {
            this.signalBuffer.shift();
        }
        const sum = this.signalBuffer.reduce((a, b) => a + b, 0);
        return sum / this.signalBuffer.length;
    }



    static applyEma(current: number, previous: number, alpha: number = 0.3): number { 
        return alpha * current + (1 - alpha) * previous;
    }

    static spring(t: number, tension: number = 50, friction: number = 7): number { 
        const s = tension / 100;
        const f = friction / 10;


        return Math.pow(Math.E, -f * t) * Math.sin(s * t) + 1;
    }



    static lerp(start: number, end: number, t: number): number { 
        return start * (1 - t) + end * t;
    }



    static clamp(val: number, min: number, max: number): number {
        return Math.max(min, Math.min(max, val));
    }



    static getVelocity(current: number, previous: number, dt: number): number { 
        return (current - previous) / dt;
    }
}
