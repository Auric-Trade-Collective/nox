#ifndef PLUGIN_H
#define PLUGIN_H

typedef struct {

} PluginCtx;

typedef void (*pluginMain)(PluginCtx*);

static inline void InvokePluginMain(PluginCtx *ctx, pluginMain cb) {
    cb(ctx);
}

#endif
