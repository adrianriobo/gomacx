#import "nsrunningapplication.h"

void* FrontmostApplication(){
    @autoreleasepool {
        return [[NSWorkspace sharedWorkspace] frontmostApplication];
        }
}

const char* BundleIdentifier(void* nsRunningApplication) {
    NSRunningApplication* a = (NSRunningApplication*)nsRunningApplication;
    pid_t pid = [a processIdentifier];
    return [[a bundleIdentifier] cStringUsingEncoding:NSISOLatin1StringEncoding];
}

CFTypeRef CreateApplicationAXRef(void* appAXRef) {
    @autoreleasepool {
        NSRunningApplication* a = (NSRunningApplication*)appAXRef;
        pid_t pid = [a processIdentifier];
        AXUIElementRef appRef = AXUIElementCreateApplication(pid);
        if (appRef == nil)
            NSLog(@"Error getting the ref app \n");
        return appRef;
    }
}

CFTypeRef GetAXFocusedWindow(CFTypeRef appAXRef) {
    @autoreleasepool {
        AXUIElementRef appRef = (AXUIElementRef)appAXRef;
        AXError err;
        AXUIElementRef focusedWindow = nil;
        err = AXUIElementCopyAttributeValue(appRef, kAXFocusedWindowAttribute,
                                        (CFTypeRef *) &focusedWindow);
        assert(kAXErrorSuccess == err);
        return focusedWindow;
    }
}
