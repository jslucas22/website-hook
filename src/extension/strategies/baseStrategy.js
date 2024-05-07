class BaseStrategy {
  constructor() {
    this._strategy = null;
  }

  set strategy(strategy) {
    this._strategy = strategy;
  }

  get strategy() {
    return this._strategy;
  }

  login(strategy) {
    return this.strategy.login(strategy)
  }
}