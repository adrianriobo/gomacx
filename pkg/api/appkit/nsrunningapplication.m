#import "nsrunningapplication.h"

void* FrontmostApplication(){
    @autoreleasepool {
        return [[NSWorkspace sharedWorkspace] frontmostApplication];
    }
}

void* FindRunningApplication(const char *bundleID, const char *windowTitle) {
    @autoreleasepool {
        NSString *nBundleID = [NSString stringWithUTF8String:bundleID];
        NSLog(@"Bundle is %@ \n", nBundleID);
        NSString *nWindowTitle = [NSString stringWithUTF8String:windowTitle];
        NSLog(@"WindowTitle is %@ \n", nWindowTitle);
        NSWorkspace* workspace = [NSWorkspace sharedWorkspace];
        NSArray* runningApps = [workspace runningApplications];
        for (NSRunningApplication* app in runningApps) {
            // pid_t pid = [app processIdentifier];
            // AXUIElementRef appRef = AXUIElementCreateApplication(pid);
            // AXError err;
            // AXUIElementRef focusedWindow = nil;
            // err = AXUIElementCopyAttributeValue(appRef, kAXFocusedWindowAttribute,
            //                             (CFTypeRef *) &focusedWindow);
            // NSString *att = nil;
            // err = AXUIElementCopyAttributeValue(focusedWindow, kAXTitleAttribute, (CFTypeRef *) &att);
            // NSString* windowTitleString = [[NSString alloc] initWithCharacters:CFStringGetCharactersPtr(att) length:CFStringGetLength(att)];
            NSLog(@"app bundle %@ \n", [app bundleIdentifier]);
            if ([[app bundleIdentifier] isEqualToString:nBundleID]) {
                NSLog(@"found\n");
                // [windowTitleString containsString:title]
                return app;
            }
        }
        return nil;
    }
}

void ShowAllApplications() {
    @autoreleasepool {
        NSWorkspace *workspace = [NSWorkspace sharedWorkspace];
        NSArray *applications = [workspace runningApplications];

        for (NSRunningApplication *app in applications) {
            NSLog(@"%@", [app bundleIdentifier]);
            pid_t pid = [app processIdentifier];
            AXUIElementRef appRef = AXUIElementCreateApplication(pid);
            AXError err;
            AXUIElementRef focusedWindow = nil;
            err = AXUIElementCopyAttributeValue(appRef, kAXFocusedWindowAttribute,
                                        (CFTypeRef *) &focusedWindow);
            NSString *att = nil;
            err = AXUIElementCopyAttributeValue(focusedWindow, kAXTitleAttribute, (CFTypeRef *) &att);
            NSLog(@"AX Tittle is %@ \n", att);
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
