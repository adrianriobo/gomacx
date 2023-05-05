#import "nsrunningapplication.h"

void* FrontmostApplication(){
    @autoreleasepool {
        return [[NSWorkspace sharedWorkspace] frontmostApplication];
        }
}

void ShowAllApplications() {
    @autoreleasepool {
        NSWorkspace *workspace = [NSWorkspace sharedWorkspace];
        NSArray *applications = [workspace runningApplications];

        for (NSRunningApplication *app in applications) {
            NSLog(@"%@", [app bundleIdentifier]);
        }
    }
}

const char* BundleIdentifier(void* nsRunningApplication) {
    NSRunningApplication* a = (NSRunningApplication*)nsRunningApplication;
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
