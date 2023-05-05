#import "axuielement.h"

// https://www.electronjs.org/docs/latest/tutorial/accessibility#macos
// CFStringRef kAXManualAccessibility = CFSTR("AXManualAccessibility");

const char* HasAXUIElementChildren(CFTypeRef axuielement) {
    @autoreleasepool {
        AXUIElementRef a = (AXUIElementRef)axuielement;
        AXError err;
        CFArrayRef childrenPtr;
        NSString *result = @"true";
        err = AXUIElementCopyAttributeValue(axuielement, kAXChildrenAttribute, (CFTypeRef *) &childrenPtr);
        if([childrenPtr count]==0){
            result = @"false";
        }
        return strdup([result UTF8String]);
    }
}

CFArrayRef GetAXUIElementChildren(CFTypeRef axuielement) {
    @autoreleasepool {
        AXUIElementRef a = (AXUIElementRef)axuielement;
        AXError err;
        CFArrayRef childrenPtr;
        err = AXUIElementCopyAttributeValue(axuielement, kAXChildrenAttribute, (CFTypeRef *) &childrenPtr);
        return childrenPtr;
    }
}

CFTypeRef GetChild(CFArrayRef children, CFIndex index) {
    @autoreleasepool {
        AXUIElementRef objectChild = CFArrayGetValueAtIndex(children, index);
        return objectChild;
    }
}

const char* GetTitleAttribute(CFTypeRef axuielement) {
    @autoreleasepool {
        AXUIElementRef a = (AXUIElementRef)axuielement;
        
        AXError err;
        NSString *att = nil;
        err = AXUIElementCopyAttributeValue(axuielement, kAXTitleAttribute, (CFTypeRef *) &att);
        if (err != kAXErrorSuccess) 
            return "";
        // NSLog(@"AX Tittle is %@ \n", att);
        return strdup([att UTF8String]);
    }
}

const char* GetRoleAttribute(CFTypeRef axuielement) {
    @autoreleasepool {
        AXUIElementRef a = (AXUIElementRef)axuielement;
        AXError err;
        NSString *att = nil;
        err = AXUIElementCopyAttributeValue(axuielement, kAXRoleAttribute, (CFTypeRef *) &att);
        if (err != kAXErrorSuccess) 
            return "";
        // NSLog(@"AX Role is %@ \n", att);
        return strdup([att UTF8String]);
    }
}

const char* GetValueAttribute(CFTypeRef axuielement) {
    @autoreleasepool {
        AXUIElementRef a = (AXUIElementRef)axuielement;
        AXError err;
        NSString *att = nil;
        err = AXUIElementCopyAttributeValue(axuielement, kAXValueAttribute, (CFTypeRef *) &att);
        if (err != kAXErrorSuccess) 
            return "";
        // NSLog(@"AX Value is %@ \n", att);
        if (CFGetTypeID(att) == CFStringGetTypeID()) {
            return strdup([att UTF8String]);
        }
        return "";
    }
}

void ClickButton(CFTypeRef axuielement) {
    @autoreleasepool {
        AXError err;
        err = AXUIElementPerformAction((AXUIElementRef)axuielement, kAXPressAction);
        if (err == kAXErrorActionUnsupported) 
            NSLog(@"error kAXErrorActionUnsupported \n");
        if (err == kAXErrorIllegalArgument) 
            NSLog(@"error kAXErrorIllegalArgument \n");
        if (err == kAXErrorInvalidUIElement) 
            NSLog(@"error kAXErrorInvalidUIElement \n");
        if (err == kAXErrorCannotComplete) 
            NSLog(@"error kAXErrorCannotComplete \n");
        if (err == kAXErrorNotImplemented) 
            NSLog(@"error kAXErrorNotImplemented \n");
    }
}

