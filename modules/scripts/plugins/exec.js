const { exec } = require('child_process');

class scriptWebpackPlugin {
  constructor(options) {
    this.options = options;
  }

  apply(compiler) {
    compiler.hooks.beforeCompile.tapPromise(
      'ScriptWebpackPlugin',
      async () => {
        const { scripts } = this.options;
        if (Array.isArray(scripts)) {
          await this.executeScriptsSequentially(scripts);
        } else if (typeof scripts === 'string') {
          await this.executeScript(scripts);
        }
      }
    );

    compiler.hooks.done.tap(
      'ScriptWebpackPlugin',
      stats => {
        const { doneCompilationMessage } = this.options;
        const message = doneCompilationMessage ?? '';
        console.log(
          '\x1b[1m\x1b[32m%s\x1b[0m',
          `[FIGOTSX] ${message}` || 'Done Compilation'
        );
      }
    );
  }

  async executeScriptsSequentially(scripts) {
    for (let i = 0; i < scripts.length; i++) {
      const script = scripts[i];
      await this.executeScript(script);
    }
  }

  executeScript(script) {
    return new Promise(resolve => {
      const { catchMessage, doneMessage } = this.options;
      const catchM = catchMessage || 'Fatal Error';
      const doneM = doneMessage || 'Complete exec Scripts';
      exec(script, (error, stdout, stderr) => {
        if (error) {
          console.log(
            '\x1b[1m\x1b[31m%s\x1b[0m',
            `[FIGOTSX] ${catchM} ${error.message} 🔴⁉️`
          );
          resolve();
          return;
        }
        if (stderr) {
          console.log(
            '\x1b[1m\x1b[33m%s\x1b[0m',
            `[FIGOTSX] ${catchM} ${stderr} 🟡⁉️`
          );
          resolve();
          return;
        }
        console.log(
          '\x1b[1m\x1b[32m%s\x1b[0m',
          `[FIGOTSX] ${doneM}`
        );
        if (stdout.length > 0)
          console.log(
            '\x1b[1m\x1b[32m%s\x1b[0m',
            `[FIGOTSX] ${stdout}`
          );
        resolve();
      });
    });
  }
}

module.exports = scriptWebpackPlugin;
