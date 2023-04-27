#import "runningapplication.h"

void* GetRunningApplications(){
    @autoreleasepool {
        return [[NSWorkspace sharedWorkspace] frontmostApplication];
        }
}

const char* GetBundleID(void* app) {
    NSRunningApplication* a = (NSRunningApplication*)app;
    pid_t pid = [a processIdentifier];
    AXUIElementRef appUI = AXUIElementCreateApplication(pid);
    AXUIElementRef frontWindow = nil;
    NSString *title = nil;
    NSString *path = nil;
    AXError err;
    // get the focused window for the application
    err = AXUIElementCopyAttributeValue(appUI, kAXFocusedWindowAttribute,
                                      (CFTypeRef *) &frontWindow);
    return [[a bundleIdentifier] cStringUsingEncoding:NSISOLatin1StringEncoding];
}
